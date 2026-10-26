package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = ((1.5 - (1.5 + 1.5)) + 1e10)
    var v3 int16 = -1
    for fi4 := 0; fi4 < 14; fi4++ {
        if false {
            if true {
                switch int(v3) {
                case 329:
                    fmt.Println(fi4, v3)
                case 127:
                    switch (fi4 % (((^(fi4))) | 1)) {
                    case -1:
                        v3 = -32768
                    case 32768:
                        v5 := GenS1{F0: 42}
                        _ = v5
                    default:
                        fmt.Println(v2)
                    }
                case -1:
                    for fi6 := 0; fi6 < 3; fi6++ {
                        var v7 uint8 = 255
                        v1 -= v1
                        _ = fi6
                        _ = v7
                    }
                default:
                    if ("hello" != "hello") {
                        v2 = (v2 / v2)
                        v1 += (^(v1))
                    }
                }
                for fi8 := 0; fi8 < 20; fi8++ {
                    v9 := GenS1{F0: int64(v1)}
                    var v10 int32 = (v1 / ((v1) | 1))
                    v11 := []int{fi4, (-(fi8))}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
            }
        } else {
            if (65535 <= 565) {
                var v12 int8 = -33
                _ = v12
            } else {
                switch (fi4 * (fi4 ^ 128)) {
                case 4294967295:
                    fmt.Println(fi4, v2)
                case 1:
                    var v13 uint = 798
                    _ = v13
                default:
                    v14 := GenS2{F0: -128}
                    _ = v14
                }
                v15 := []int{}
                _ = v15
            }
            v16 := []int{fi4, fi4, 2147483647, fi4}
            var v17 string = "foo bar"
            _ = v16
            _ = v17
        }
        v3 = v3
        v18 := GenS2{F0: 1}
        _ = fi4
        _ = v18
    }
    v19 := GenS2{F0: 0}
    if true {
        v1 = v1
    }
    switch ("\n\t" + "\n\t") {
    case "":
        fmt.Println(v2)
    }
    var v20 int64 = 0
    switch "a" {
    case "":
        var v21 int64 = v20
        _ = v21
    default:
        v20 = (v20 | (v20 % (((v20 & (v20 + 1))) | 1)))
    }
    v20 = (v20 >> ((65535) & 63))
    v20 += v20
    switch ("foo bar" + "") {
    case "hello":
        fmt.Println(v1)
    case "":
        for fi22 := 0; fi22 < 20; fi22++ {
            if false {
                v23 := []int{32768, -9223372036854775808, fi22}
                v3 = 0
                for fi24 := 0; fi24 < 4; fi24++ {
                    v25 := GenS2{F0: -1}
                    fmt.Println(v3)
                    v1 += v1
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                var v26 uint64 = 390
                v27 := []int{127, fi22, fi22}
                _ = v26
                _ = v27
            }
            _ = fi22
        }
    }
    v28 := GenS3{F0: (^(v3)), F1: 65535, F2: 460}
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS4{F0: 1}
        for fi31 := 0; fi31 < 18; fi31++ {
            switch fi31 {
            case 32767:
                v20 = v20
            default:
                v20 += v20
            }
            v1 = (v1 - (v1 << ((777) & 63)))
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    v32 := []int{}
    v33 := GenS1{F0: v20}
    var v34 string = ""
    v35 := GenS2{F0: 0}
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    v3 = (v3 | (v3 >> ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
