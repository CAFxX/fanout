package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint64, p1 int32, p2 uint8) uint8 {
    if false {
        var v5 int = 0
        v6 := []int{}
        if (!(-1 == 1)) {
            fmt.Println(p2)
            var v7 uint = 0
            var v8 float32 = float32(-2.25)
            _ = v7
            _ = v8
        }
        _ = v5
        _ = v6
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint64 = p0
            if (-128 <= 41) {
                switch fi9 {
                case -2147483648:
                    var v11 uint16 = 42
                    _ = v11
                }
                if (p1 > 0) {
                    var v12 int16 = -1
                    _ = v12
                } else {
                    p0 += (v10 % (((v10 % (((p0 >> ((255) & 63))) | 1))) | 1))
                    switch "" {
                    case "foo bar":
                        var v13 float64 = 1.5
                        _ = v13
                    case "a":
                        p1 = p1
                    case "":
                        var v14 uint16 = uint16(p2)
                        _ = v14
                    }
                }
            }
            _ = fi9
            _ = v10
        }
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS1{F0: uint(p2)}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 uint, p1 uint32, p2 uint64) int64 {
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            switch "a" {
            case "":
                var v19 uint32 = 0
                _ = v19
            default:
                v20 := GenS2{F0: p1, F1: p1, F2: ""}
                _ = v20
            }
            _ = fi18
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint32 = 4294967295
            _ = fi21
            _ = v22
        }
        switch 32768 {
        case 128:
            if (p2 > (p2 % ((p2) | 1))) {
                v23 := GenS2{F0: 2147483647, F1: p1, F2: "\n\t"}
                p1 += (^(p1))
                p1 = (p1 >> ((340) & 63))
                _ = v23
            }
        case 32768:
            fmt.Println(p1, fi17)
        case 42:
            p2 = (^(p2))
        default:
            var v24 uint16 = 0
            _ = v24
        }
        _ = fi17
    }
    p1 = (p1 << (((p0 * p0)) & 63))
    p0 += 18446744073709551615
    p1 += p1
    if true {
        v25 := []int{127, -65, 42, 0}
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint8 = 1
    v4 := []uint32{42, 434, 2147483648}
    v26 := GenS2{F0: 1, F1: 65535, F2: "hello"}
    fmt.Println(v3)
    for fi27 := 0; fi27 < 2; fi27++ {
        v3 = (v3 * 231)
        if (0 >= 572) {
            var v28 int = fi27
            _ = v28
        } else {
            v29 := GenS1{F0: 42}
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 int64 = 42
                v3 -= 255
                v3 -= genFn0(0, 681, (v3 * (v3 ^ (v3 >> ((334) & 63)))))
                _ = fi30
                _ = v31
            }
            _ = v29
        }
        _ = fi27
    }
    var v32 int = -9223372036854775808
    if false {
        v3 = (v3 | 0)
        delete(v1, "")
    }
    v32 = v32
    for fi33 := 0; fi33 < 4; fi33++ {
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        v34 := GenS2{F0: 2147483647, F1: 0, F2: ((("a" + "hello") + (("" + "a") + "\n\t")) + "\n\t")}
        _ = fi33
        _ = v34
    }
    if true {
        v35 := v2[(("hello" + "hello") + "\n\t")]
        _ = v35
    } else {
        for fi36 := 0; fi36 < 6; fi36++ {
            fmt.Println(v32)
            _ = fi36
        }
        v3 = v3
    }
    var v37 int64 = 1
    v38 := GenS3{F0: -1}
    if (!(false || ("" == "hello"))) {
        v39 := v2["foo bar"]
        v40 := GenS4{F0: 42, F1: 30}
        v41 := GenS4{F0: -181, F1: v3}
        _ = v39
        _ = v40
        _ = v41
    } else {
        v3 = (v3 << ((uint(v3)) & 63))
    }
    switch "a" {
    case "hello":
        for fi42 := 0; fi42 < 9; fi42++ {
            _ = len(v4)
            var v43 int16 = 132
            switch ("foo bar" + ("foo bar" + (("foo bar" + ("\n\t" + ("a" + ""))) + ""))) {
            case "\n\t":
                if true {
                    if false {
                        if len(v4) > 0 {
                            v4[0] = 937
                        }
                    } else {
                        v32 = 159
                    }
                    v44 := GenS2{F0: uint32(v3), F1: 722, F2: "foo bar"}
                    _ = v44
                }
            }
            _ = fi42
            _ = v43
        }
    case "a":
        var v45 uint64 = 0
        _ = v45
    default:
        for fi46 := 0; fi46 < 15; fi46++ {
            fmt.Println(v32, fi46, v32)
            for fi47 := 0; fi47 < 18; fi47++ {
                v48 := GenS1{F0: 1}
                v49 := GenS3{F0: -1}
                switch (fi47 * (fi47 % (((fi47 << ((18446744073709551615) & 63))) | 1))) {
                case 42:
                    var v50 uint64 = 18446744073709551615
                    _ = v50
                }
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = fi46
        }
    }
    for fi51 := 0; fi51 < 9; fi51++ {
        if (!(("a" >= (("a" + "a") + (("\n\t" + (("hello" + "foo bar") + "foo bar")) + ("a" + ("hello" + "foo bar"))))) || (2 < v3))) {
            v3 = genFn0(1000, 1, 42)
        } else {
            for fi52 := 0; fi52 < 19; fi52++ {
                var v53 uint32 = 2
                _ = fi52
                _ = v53
            }
            var v54 uint32 = 2147483647
            _ = v54
        }
        _ = fi51
    }
    var v55 uint8 = (v3 * genFn0(2, 2147483647, v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v55
}
