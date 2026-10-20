package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 6; fi4++ {
        _ = len(v1)
        _ = fi4
    }
    if (32767 <= -347) {
        if false {
            var v5 uint16 = 527
            var v6 int16 = -1
            var v7 uint64 = 656
            _ = v5
            _ = v6
            _ = v7
        } else {
            if ((true || false) && ((318 < 395) && true)) {
                var v8 bool = true
                _ = len(v1)
                _ = v8
            } else {
                v9 := GenS1{F0: true}
                _ = v9
            }
            var v10 int = -2147483648
            _ = v10
        }
        v11 := GenS2{F0: 18446744073709551615, F1: 42}
        _ = v11
    } else {
        v12 := []int{42, 2147483647}
        v12 = append(v12, 1)
        _ = v12
    }
    var v13 uint = 704
    var v14 int16 = 32767
    if ((2 <= 42) || true) {
        var v15 float64 = float64(float32((1.5 / 3.14159)))
        if (false || true) {
            var v16 float32 = 0.0
            v17 := GenS2{F0: 997, F1: 0}
            v14 = 42
            _ = v16
            _ = v17
        }
        v15 -= v15
        _ = v15
    }
    switch ((("\n\t" + "a") + ("hello" + "foo bar")) + "a") {
    case "a":
        var v18 int16 = -1
        _ = v18
    case "hello":
        v19 := []int{}
        _ = v19
    case "":
        for fi20 := 0; fi20 < 7; fi20++ {
            switch fi20 {
            case 4294967295:
                v14 = 32767
            }
            v14 -= v14
            _ = fi20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 float32 = 0.0
        if false {
            v23 := []int{}
            if false {
                var v24 int64 = -902
                v22 -= (v22 + v22)
                var v25 int16 = (v14 % (((v14 + -940)) | 1))
                _ = v24
                _ = v25
            }
            _ = v23
        }
        var v26 string = "foo bar"
        _ = fi21
        _ = v22
        _ = v26
    }
    var v27 int8 = -24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
}
