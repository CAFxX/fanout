package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
func genFn0(p0 int16) int {
    fmt.Println(p0)
    v4 := []int{}
    fmt.Println(p0)
    v5 := GenS1{F0: true, F1: -1}
    _ = p0
    _ = v4
    _ = v5
    return -611
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    var v6 bool = (255 != 149)
    var v7 uint64 = 1
    v8 := GenS1{F0: ((v7 >> ((uint(v7)) & 63)) >= (v7 / ((v7) | 1))), F1: -2}
    var v9 int = 1
    _ = len(v1)
    v3[("hello" + "foo bar")] = (0.0 < (((1e10 * (float64(1.5) * 0.1)) * -2.25) * -2.25))
    var v10 int8 = 113
    switch ((("a" + "a") + ("foo bar" + "a")) + ("\n\t" + "a")) {
    case "hello":
        v9 = 32768
    case "a":
        switch (v9 ^ v9) {
        case -732:
            v6 = (float64(((-2.25 * (3.14159 + 0.0)) + 0.1)) < 3.14159)
        }
    }
    fmt.Println(v9, v6, v9)
    var v11 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
