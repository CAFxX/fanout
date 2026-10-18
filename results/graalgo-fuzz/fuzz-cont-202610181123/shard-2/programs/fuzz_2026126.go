package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int16 }
func genFn0() string {
    v4 := []int{}
    var v5 int16 = 1
    var v6 int16 = (^(v5))
    _ = v4
    _ = v5
    _ = v6
    return (("hello" + "foo bar") + "hello")
}
func genFn1(p0 int) uint64 {
    v7 := GenS1{F0: 242, F1: (!true), F2: p0}
    switch (("hello" + "foo bar") + "hello") {
    case "foo bar":
        fmt.Println(p0)
    case "hello":
        if true {
            p0 -= (p0 | 42)
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 = -9223372036854775808
                v9 := []int{fi8, -2147483648}
                switch 128 {
                case -779:
                    p0 = fi8
                }
                _ = fi8
                _ = v9
            }
        }
    case "a":
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS2{F0: ""}
            _ = fi10
            _ = v11
        }
    default:
        v12 := []int{p0, 245, (p0 / (((p0 | p0)) | 1))}
        _ = v12
    }
    var v13 string = "foo bar"
    v13 = "a"
    v14 := GenS3{F0: -1, F1: p0}
    _ = p0
    _ = v13
    _ = v14
    _ = v7
    return 2
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v15 uint64 = 2
    if (("" < "foo bar") || false) {
        for fi16 := 0; fi16 < 1; fi16++ {
            fmt.Println(fi16)
            fmt.Println(v15, v15, v1)
            var v17 int32 = -2147483648
            _ = fi16
            _ = v17
        }
        if true {
            v18 := GenS1{F0: 255, F1: (941 == 65535), F2: 1}
            if ((!(!true)) && true) {
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 int32 = 970
                    _ = fi19
                    _ = v20
                }
            } else {
                v15 = genFn1(1)
                var v21 uint32 = 0
                v22 := GenS4{F0: 32767, F1: v15}
                _ = v21
                _ = v22
            }
            v23 := GenS4{F0: 42, F1: 42}
            _ = v18
            _ = v23
        }
    } else {
        switch 0 {
        case 42:
            for fi24 := 0; fi24 < 15; fi24++ {
                v1 = ((float64(1e10) - v1) - (v1 / (v1 * v1)))
                _ = fi24
            }
        case 4294967295:
            var v25 uint8 = 1
            _ = v25
        default:
            v1 = (0.0 + (float64(0.1) / float64(1e10)))
        }
    }
    var v26 bool = (!((127 < 204) && true))
    v27 := []int{}
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 bool = (9223372036854775807 >= 1)
        _ = fi28
        _ = v29
    }
    _ = len(v3)
    v30 := GenS3{F0: 2147483647, F1: 32768}
    v31 := GenS2{F0: "a"}
    v26 = (82 > 127)
    switch "foo bar" {
    case "foo bar":
        v3[0] = v26
    }
    var v32 uint = 0
    v26 = (-25 <= -5)
    _ = len(v2)
    if true {
        v32 = v32
        var v33 int8 = -1
        _ = v33
    } else {
        switch 984 {
        case 32767:
            v1 += 0.0
        default:
            v34 := GenS1{F0: 255, F1: (!true), F2: -160}
            _ = v34
        }
        if (1 == 115) {
            for fi35 := 0; fi35 < 5; fi35++ {
                for fi36 := 0; fi36 < 10; fi36++ {
                    var v37 uint16 = 610
                    v38 := GenS3{F0: -462, F1: 42}
                    fmt.Println(fi35, v1, v37)
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = fi35
            }
            delete(v3, 231)
            v39 := GenS3{F0: -1, F1: 32767}
            _ = v39
        }
    }
    v40 := GenS5{F0: 792}
    if len(v27) > 0 {
        v27[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v40
}
