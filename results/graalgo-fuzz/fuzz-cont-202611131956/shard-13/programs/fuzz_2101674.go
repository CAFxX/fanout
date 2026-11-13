package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 string }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -474
    v3 := []int{}
    switch (("\n\t" + "") + "") {
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        fmt.Println(v2)
    }
    v4 := GenS1{F0: false, F1: 99, F2: (^(v2))}
    if (true && (!false)) {
        delete(v1, "hello")
        var v5 int16 = -110
        var v6 int8 = -1
        _ = v5
        _ = v6
    }
    switch 42 {
    case 834:
        if len(v3) > 0 {
            v3[0] = 42
        }
    default:
        var v7 int64 = -1
        _ = v7
    }
    v8 := GenS2{F0: 128, F1: 2}
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v2 = v2
            v2 -= 2147483647
            fmt.Println(fi9)
            _ = fi10
        }
        switch (fi9 ^ fi9) {
        case 563:
            if false {
                var v11 int = fi9
                _ = v11
            }
        case 0:
            if false {
                if true {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v2 -= v2
                        _ = fi12
                    }
                    v13 := GenS2{F0: 42, F1: 827}
                    _ = v13
                } else {
                    v2 += (v2 * (v2 ^ v2))
                    for fi14 := 0; fi14 < 12; fi14++ {
                        var v15 uint64 = 664
                        _ = fi14
                        _ = v15
                    }
                }
            } else {
                v2 -= v2
                v2 += v2
                var v16 float32 = 3.14159
                _ = v16
            }
        case 4294967295:
            var v17 string = "a"
            _ = v17
        default:
            v2 = (v2 ^ 0)
        }
        _ = fi9
    }
    v18 := GenS3{F0: true, F1: 127, F2: "\n\t"}
    switch -9223372036854775808 {
    case 13:
        switch ("\n\t" + ("hello" + "")) {
        case "\n\t":
            v2 -= 42
        default:
            v19 := GenS1{F0: false, F1: 4294967295, F2: (^(v2))}
            _ = v19
        }
    }
    if false {
        v2 -= (v2 << ((960) & 63))
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 uint8 = 0
            _ = fi20
            _ = v21
        }
        v22 := v1["hello"]
        _ = v22
    }
    fmt.Println(v2)
    switch 42 {
    case -2147483648:
        v1["a"] = false
    default:
        v23 := GenS1{F0: true, F1: 758, F2: v2}
        _ = v23
    }
    switch "hello" {
    case "":
        v2 = v2
    case "hello":
        if (1 != 1000) {
            _ = len(v3)
        }
    case "\n\t":
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int16 = 1000
            v26 := GenS1{F0: (true || true), F1: 332, F2: -254}
            v27 := GenS3{F0: true, F1: 100, F2: "foo bar"}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    }
    _ = len(v1)
    switch "a" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        fmt.Println(v2)
    case "":
        v28 := GenS4{F0: v2}
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
