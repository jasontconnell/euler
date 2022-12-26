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

$contents = Get-Content -Path ".\boilerplate\solve.txt"

$alphanum = $puzzle.PadLeft(4-$puzzle.Length, "0")
$contents = $contents -replace "#0#", $puzzle
$contents = $contents -replace "#1#", $alphanum

Set-Content -Path ".\puzzles\$alphanum.go" -Value $contents
