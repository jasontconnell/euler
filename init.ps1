param (
    [string] $puzzle
)

if ($puzzle -eq "") {
    Write-Host "Need puzzle number"
    return
}

if (-not (Test-Path -Path ".\puzzles")) {
    New-Item -Path ".\puzzles" -ItemType "directory"
}

$alphanum = $puzzle.PadLeft(3, "0")
if (-not (Test-Path -Path ".\puzzles\$alphanum.go")) {

    $contents = Get-Content -Path ".\boilerplate\solve.txt"
    $contents = $contents -replace "#0#", $puzzle
    $contents = $contents -replace "#1#", $alphanum

    Set-Content -Path ".\puzzles\$alphanum.go" -Value $contents
}