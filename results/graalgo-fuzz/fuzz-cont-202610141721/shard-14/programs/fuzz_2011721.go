package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
func genFn0() uint {
    var v4 float32 = 0.1
    fmt.Println(v4)
    v5 := GenS1{F0: "\n\t", F1: (("\n\t" + (("foo bar" + "foo bar") + "")) + ("hello" + "\n\t"))}
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -32768
    var v3 bool = false
    v2 = v2
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6, v3)
        if false {
            v7 := GenS1{F0: "\n\t", F1: ""}
            var v8 int64 = 0
            for fi9 := 0; fi9 < 4; fi9++ {
                v8 = -777
                _ = fi9
            }
            _ = v7
            _ = v8
        } else {
            switch "a" {
            case "a":
                v10 := GenS1{F0: "\n\t", F1: "hello"}
                _ = v10
            }
            v11 := GenS2{F0: fi6}
            v3 = false
            _ = v11
        }
        _ = fi6
    }
    v12 := GenS2{F0: -579}
    var v13 uint64 = 758
    v3 = ((v3 || ((!true) && (-1 <= -1))) || true)
    if ((1.5 > 1e10) && v3) {
        var v14 float64 = (3.14159 - 1e10)
        _ = v14
    } else {
        var v15 int32 = 1
        _ = v15
    }
    switch -1 {
    case 32768:
        v13 = v13
    }
    if (1 == int64(v13)) {
        for fi16 := 0; fi16 < 14; fi16++ {
            if (!(((!(v2 < (v2 ^ v2))) || (3.14159 != (0.1 + 3.14159))) && v3)) {
                delete(v1, "\n\t")
                var v17 int16 = v2
                _ = v17
            }
            var v18 string = "\n\t"
            v19 := []int{}
            _ = fi16
            _ = v18
            _ = v19
        }
    } else {
        _ = len(v1)
        fmt.Println(v3, v2)
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint8 = 255
        v22 := GenS3{F0: 42}
        _ = fi20
        _ = v21
        _ = v22
    }
    if (0 != 111) {
        v1["hello"] = v3
        switch "\n\t" {
        case "hello":
            v23 := []int{-2147483648, -1, 32767, 2147483647}
            _ = v23
        default:
            var v24 string = ""
            _ = v24
        }
    } else {
        var v25 int = 32767
        var v26 float64 = (1.5 * 1e10)
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
