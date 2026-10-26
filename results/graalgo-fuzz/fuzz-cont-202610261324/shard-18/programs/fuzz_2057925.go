package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32, p1 int8, p2 uint32) int16 {
    var v6 uint64 = 42
    v7 := GenS1{F0: 53}
    switch "hello" {
    case "":
        p0 = p0
    case "foo bar":
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 825
    var v2 float32 = 3.14159
    var v3 uint64 = 0
    var v4 uint64 = 830
    var v5 int16 = -32768
    fmt.Println(v5, v5, v5)
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int16 = 1
        _ = fi8
        _ = v9
    }
    switch 42 {
    case 606:
        var v10 int = 737
        _ = v10
    default:
        for fi11 := 0; fi11 < 4; fi11++ {
            v4 += 393
            switch 0 {
            case 4294967295:
                if (!false) {
                    v12 := []int{128, (fi11 * fi11), (-(fi11)), int(v5)}
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v5 = 1
                        v14 := GenS2{F0: 56, F1: 65535}
                        _ = fi13
                        _ = v14
                    }
                    _ = len(v12)
                    _ = v12
                } else {
                    var v15 float32 = 0.0
                    _ = v15
                }
            case -365:
                v2 = 1.5
            }
            v3 = v3
            _ = fi11
        }
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v5 = (v5 % (((-(v5))) | 1))
            var v18 float32 = v2
            var v19 uint8 = 162
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 uint32 = uint32(v4)
        _ = fi16
        _ = v20
    }
    var v21 float32 = v2
    for fi22 := 0; fi22 < 9; fi22++ {
        v21 += (v2 + 1.5)
        _ = fi22
    }
    v2 -= 1e10
    switch "foo bar" {
    case "a":
        var v23 int64 = 503
        _ = v23
    case "\n\t":
        v24 := GenS3{F0: "\n\t", F1: -34}
        _ = v24
    }
    v25 := []int{-406}
    v26 := GenS3{F0: ((((("" + "a") + "") + ("foo bar" + ("a" + "a"))) + "foo bar") + (("\n\t" + "hello") + "a")), F1: 110}
    if ((95 == -1) && true) {
        var v27 float64 = (((3.14159 + 1e10) * 1.5) * float64(((v21 * 0.0) - (v21 + v2))))
        v21 = v21
        if (-1 <= 100) {
            var v28 bool = true
            _ = v28
        } else {
            v29 := GenS3{F0: "hello", F1: -26}
            if true {
                if true {
                    if false {
                        var v30 uint32 = 566
                        _ = v30
                    }
                    if (65535 == 412) {
                        v4 += (v3 - (v4 * v4))
                        _ = len(v25)
                    }
                    v31 := GenS2{F0: 170, F1: 2}
                    _ = v31
                } else {
                    fmt.Println(v5, v2)
                    v3 = v4
                }
                switch 2147483647 {
                case 176:
                    for fi32 := 0; fi32 < 19; fi32++ {
                        v21 = v2
                        _ = fi32
                    }
                case 0:
                    switch "" {
                    case "a":
                        fmt.Println(v4, v1, v5)
                    default:
                        var v33 string = "foo bar"
                        _ = v33
                    }
                default:
                    var v34 int8 = 95
                    _ = v34
                }
            } else {
                var v35 uint8 = 128
                v21 -= 1e10
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = v35
            }
            _ = v29
        }
        _ = v27
    }
    v36 := GenS4{F0: v5}
    v21 = (3.14159 / 1e10)
    switch (("foo bar" + ("\n\t" + "a")) + "hello") {
    case "\n\t":
        if true {
            v2 = 1e10
        }
    case "foo bar":
        v3 = 42
    }
    fmt.Println(v4, v4)
    var v37 int64 = 95
    var v38 string = "a"
    v37 = (v37 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
