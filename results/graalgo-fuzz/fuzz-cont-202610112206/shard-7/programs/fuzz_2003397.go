package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int16) uint64 {
    for fi5 := 0; fi5 < 13; fi5++ {
        if false {
            v6 := GenS1{F0: true, F1: 100}
            if ((p0 / ((42) | 1)) != (p0 - (p0 & (p0 - (p0 ^ 381))))) {
                p0 = (p0 & p0)
                v7 := GenS2{F0: 176, F1: "foo bar"}
                _ = v7
            } else {
                var v8 int = fi5
                switch "foo bar" {
                case "foo bar":
                    if true {
                        for fi9 := 0; fi9 < 1; fi9++ {
                            v10 := []int{(v8 ^ fi9), fi5}
                            _ = fi9
                            _ = v10
                        }
                        var v11 string = ((("hello" + "\n\t") + (("" + ("foo bar" + "a")) + "a")) + ("" + ("\n\t" + "hello")))
                        v12 := GenS3{F0: 0, F1: 927}
                        _ = v11
                        _ = v12
                    }
                }
                _ = v8
            }
            p0 = (p0 << ((65535) & 63))
            _ = v6
        }
        _ = fi5
    }
    if ((!false) || true) {
        v13 := []int{}
        v14 := GenS1{F0: true, F1: 1}
        _ = v13
        _ = v14
    } else {
        v15 := GenS2{F0: 2, F1: "\n\t"}
        p0 -= 32767
        _ = v15
    }
    v16 := []int{}
    _ = p0
    _ = v16
    return 42
}
func genFn1(p0 uint8, p1 int16, p2 string) uint {
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(p2, p2)
        _ = fi17
    }
    v18 := GenS1{F0: true, F1: -128}
    for fi19 := 0; fi19 < 8; fi19++ {
        p0 += (p0 << ((18446744073709551615) & 63))
        v20 := []int{fi19, (fi19 * fi19), fi19}
        var v21 int32 = -857
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint16 = 1
            var v25 uint = 18446744073709551615
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    var v26 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v26
    return 42
}

func main() {
    var v1 int8 = 63
    var v2 uint = 0
    var v3 int16 = 1
    var v4 int32 = 947
    v2 += v2
    var v27 float64 = (-2.25 + (((-2.25 / -2.25) * 0.0) / -2.25))
    v4 += -620
    v4 = v4
    var v28 bool = false
    v29 := GenS1{F0: v28, F1: v1}
    for fi30 := 0; fi30 < 5; fi30++ {
        v4 = (v4 | (v4 + (-(v4))))
        switch (("a" + "\n\t") + ((("hello" + "") + (((("a" + "a") + "") + (("" + ("" + ("a" + ("foo bar" + "foo bar")))) + "a")) + "hello")) + "")) {
        case "a":
            switch int(v2) {
            case -116:
                for fi31 := 0; fi31 < 4; fi31++ {
                    var v32 int32 = 1
                    var v33 bool = v28
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            case 32767:
                var v34 uint16 = 42
                _ = v34
            }
        case "foo bar":
            v35 := GenS3{F0: genFn0((v3 / ((v3) | 1))), F1: 4294967295}
            _ = v35
        case "hello":
            var v36 uint64 = 0
            _ = v36
        default:
            var v37 uint = v2
            _ = v37
        }
        _ = fi30
    }
    v4 = v4
    v38 := []int{42, 128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
}
