package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 int8 = 31
    var v2 float64 = float64(3.14159)
    var v3 bool = false
    var v4 bool = v3
    v4 = false
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{(fi5 | fi5)}
        v3 = v3
        fmt.Println(v4, v1)
        _ = fi5
        _ = v6
    }
    if v4 {
        var v7 uint16 = 1
        _ = v7
    } else {
        v8 := []int{-1, 2147483647}
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v8) > 0 {
                v8[0] = fi9
            }
            _ = fi9
        }
        _ = v8
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 bool = ((v4 && true) || (!v3))
        _ = fi10
        _ = v11
    }
    if v4 {
        switch ("a" + "a") {
        case "":
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int64 = -766
                _ = fi12
                _ = v13
            }
        case "hello":
            v14 := GenS1{F0: -249, F1: 779}
            _ = v14
        }
        var v15 uint = 65535
        switch -9223372036854775808 {
        case -464:
            v1 = v1
        }
        _ = v15
    } else {
        if (v3 || v3) {
            switch -1 {
            case 32768:
                v3 = (((float32((0.1 + (0.1 * (v2 - 1.5)))) > 0.1) || (983 >= 0)) || (v1 == v1))
            }
            var v16 uint = 495
            fmt.Println(v1, v3)
            _ = v16
        } else {
            var v17 uint = 255
            if v3 {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := []int{fi18, fi18}
                    _ = fi18
                    _ = v19
                }
                var v20 float64 = ((0.1 * 0.1) / v2)
                var v21 int = 32768
                _ = v20
                _ = v21
            } else {
                var v22 int16 = 42
                var v23 int8 = 127
                _ = v22
                _ = v23
            }
            v4 = (127 < 32767)
            _ = v17
        }
        v2 += v2
        var v24 uint64 = 0
        _ = v24
    }
    if v3 {
        v25 := []int{9223372036854775807}
        var v26 string = "foo bar"
        _ = v25
        _ = v26
    }
    v1 += v1
    v27 := GenS1{F0: 32768, F1: 2}
    var v28 uint8 = 1
    if true {
        v1 = v1
    } else {
        var v29 int64 = 1
        var v30 uint = 898
        _ = v29
        _ = v30
    }
    v31 := []int{2147483647, 128}
    if len(v31) > 0 {
        v31[0] = 1
    }
    switch 128 {
    case 4294967295:
        v32 := GenS1{F0: 32768, F1: 1}
        _ = v32
    case 2147483647:
        v28 = 2
    default:
        v33 := GenS2{F0: ("\n\t" + "a"), F1: 4294967295}
        _ = v33
    }
    v3 = ((float64(1.5) / v2) != float64((3.14159 + 0.1)))
    v1 = v1
    switch -284 {
    case -1:
        var v34 uint8 = v28
        _ = v34
    case -292:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
