package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8 }
func genFn0() string {
    for fi4 := 0; fi4 < 9; fi4++ {
        switch ("\n\t" + "hello") {
        case "hello":
            var v5 int32 = 42
            _ = v5
        case "\n\t":
            var v6 int8 = -128
            _ = v6
        default:
            v7 := []int{-1, 2147483647, 127}
            _ = v7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(fi4, fi8)
            _ = fi8
        }
        if (false && (((^(fi4)) <= (^(fi4))) || true)) {
            var v9 int16 = 1000
            var v10 bool = true
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    if (-569 != -200) {
        switch ("hello" + "foo bar") {
        case "\n\t":
            v11 := GenS1{F0: 355, F1: -705}
            _ = v11
        case "":
            var v12 string = ("a" + ("hello" + (("\n\t" + "foo bar") + "")))
            _ = v12
        case "a":
            v13 := GenS2{F0: 1000, F1: 42, F2: 848}
            _ = v13
        }
    }
    var v14 int = -853
    _ = v14
    return "hello"
}

func main() {
    var v1 uint64 = 42
    var v2 int16 = -1
    var v3 uint64 = v1
    v1 = 577
    v2 = v2
    v1 += (^(v1))
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v1, v1, v1)
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := []int{fi16, 32767, fi16}
            v3 = 1
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v3 = (v3 / (((v1 | v1)) | 1))
    fmt.Println(v1, v1)
    v18 := GenS1{F0: -32768, F1: 0}
    var v19 int64 = -72
    var v20 int = 0
    v1 = (^(v3))
    fmt.Println(v20)
    v19 = v19
    v21 := GenS2{F0: v2, F1: (-(v2)), F2: 0}
    v22 := GenS3{F0: 1}
    switch (^(v20)) {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 10; fi23++ {
            if (-1 != -1) {
                var v24 string = ""
                for fi25 := 0; fi25 < 18; fi25++ {
                    if (42 == 42) {
                        v1 -= v3
                        var v26 float32 = float32(-2.25)
                        _ = v26
                    }
                    var v27 int64 = v19
                    var v28 uint16 = 32767
                    _ = fi25
                    _ = v27
                    _ = v28
                }
                _ = v24
            } else {
                switch 20 {
                case 42:
                    if false {
                        switch "" {
                        case "a":
                            v19 = 42
                        }
                        fmt.Println(v20, v1, v19)
                        for fi29 := 0; fi29 < 15; fi29++ {
                            fmt.Println(fi29, v20, v2)
                            v3 = (^(v3))
                            var v30 uint = 65535
                            _ = fi29
                            _ = v30
                        }
                    } else {
                        v2 = 890
                        for fi31 := 0; fi31 < 9; fi31++ {
                            var v32 uint64 = 901
                            _ = fi31
                            _ = v32
                        }
                        var v33 uint32 = 10
                        _ = v33
                    }
                case 1:
                    v1 += 18446744073709551615
                }
                fmt.Println(v3)
            }
            _ = fi23
        }
    case 32768:
        v20 += v20
    default:
        var v34 uint8 = 255
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
