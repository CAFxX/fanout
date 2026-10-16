package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int16 }
func genFn0(p0 uint64) int64 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 255, F1: 65535}
        if false {
            var v6 uint32 = 0
            v7 := []int{fi4}
            _ = v6
            _ = v7
        } else {
            switch (fi4 ^ fi4) {
            case 1:
                v8 := GenS1{F0: 42, F1: 744}
                _ = v8
            case -9223372036854775808:
                p0 = 720
            case 32768:
                var v9 int64 = -9223372036854775808
                _ = v9
            default:
                for fi10 := 0; fi10 < 10; fi10++ {
                    p0 = (p0 + p0)
                    _ = fi10
                }
            }
        }
        _ = fi4
        _ = v5
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            p0 = p0
            var v13 uint8 = 2
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 bool = (2 >= 0)
    _ = p0
    _ = v14
    return -1
}
func genFn1() uint32 {
    switch "hello" {
    case "hello":
        switch 127 {
        case 42:
            switch ("hello" + "foo bar") {
            case "":
                var v15 string = (("a" + "foo bar") + (("\n\t" + "\n\t") + ""))
                _ = v15
            }
        case 9223372036854775807:
            v16 := []int{127, -9223372036854775808}
            _ = v16
        case 4294967295:
            var v17 string = ((("\n\t" + "") + "hello") + ("foo bar" + "\n\t"))
            _ = v17
        }
    case "":
        var v18 uint32 = 2147483647
        _ = v18
    default:
        var v19 int64 = -1
        _ = v19
    }
    var v20 uint8 = 128
    for fi21 := 0; fi21 < 19; fi21++ {
        v20 -= v20
        for fi22 := 0; fi22 < 4; fi22++ {
            v20 = 0
            v20 += (^(v20))
            var v23 int64 = 1
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 bool = false
    var v25 int8 = 127
    _ = v20
    _ = v24
    _ = v25
    return 2147483648
}

func main() {
    var v1 uint = 255
    var v2 bool = ((0 == 127) && false)
    var v3 uint64 = 418
    v3 = v3
    fmt.Println(v1, v1)
    var v26 uint64 = (v3 << (((v1 ^ v1)) & 63))
    if (v2 || v2) {
        v26 -= (v26 | v26)
        v2 = (!v2)
        v26 += v3
    }
    v27 := GenS1{F0: 218, F1: 2}
    fmt.Println(v2, v26, v26)
    v3 = v26
    var v28 int16 = 1000
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := []int{fi29, 32767}
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 255, F1: 2}
    v32 := []int{1}
    for fi33 := 0; fi33 < 12; fi33++ {
        if (-514 > 471) {
            v28 -= (v28 - (v28 & (v28 - (v28 / ((v28) | 1)))))
            v2 = false
            v1 = (^(v1))
        } else {
            _ = len(v32)
            v26 += (v3 / ((438) | 1))
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            v1 = 65535
            var v35 float32 = 1e10
            switch fi34 {
            case 2147483647:
                var v36 int = -2147483648
                _ = v36
            case 265:
                var v37 uint64 = v3
                _ = v37
            case 1:
                v38 := GenS2{F0: -32768, F1: (fi34 / (((fi34 + fi33)) | 1)), F2: 2147483647}
                _ = v38
            default:
                v39 := GenS3{F0: 113}
                _ = v39
            }
            var v40 uint = 65535
            _ = fi34
            _ = v35
            _ = v40
        }
        var v41 float32 = ((0.0 + (3.14159 / 1.0)) * 0.1)
        _ = fi33
        _ = v41
    }
    v42 := GenS4{F0: (-(v28)), F1: v28}
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 int64 = 1
        fmt.Println(v3)
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v42
}
