package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint8) uint32 {
    if (!true) {
        if (0.1 <= ((1.5 - 0.1) - -2.25)) {
            var v7 uint8 = 42
            v8 := GenS1{F0: 0, F1: 515}
            _ = v7
            _ = v8
        } else {
            var v9 uint32 = 1
            v10 := []int{4294967295, 32767, 290}
            if (-2147483648 > -641) {
                var v11 string = ((("\n\t" + (("hello" + (("hello" + "a") + "a")) + "foo bar")) + "") + "hello")
                for fi12 := 0; fi12 < 9; fi12++ {
                    v11 = (v11 + v11)
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi12
                }
                _ = v11
            } else {
                p0 += (^(p1))
                if (!(!false)) {
                    v13 := GenS1{F0: 4294967295, F1: 36}
                    _ = v13
                } else {
                    var v14 uint32 = v9
                    _ = v14
                }
                v15 := GenS2{F0: 0}
                _ = v15
            }
            _ = v10
            _ = v9
        }
    } else {
        p1 = p0
        p1 = p1
        fmt.Println(p1, p0)
    }
    var v16 int64 = 1
    switch 32767 {
    case 128:
        var v17 uint8 = p1
        _ = v17
    case 5:
        v18 := []int{0}
        _ = v18
    case 4294967295:
        var v19 uint32 = 1
        _ = v19
    default:
        p0 -= (p1 << ((978) & 63))
    }
    _ = p0
    _ = p1
    _ = v16
    return 1
}
func genFn1(p0 uint, p1 uint64, p2 int) int8 {
    p1 = (^(p1))
    p1 = p1
    fmt.Println(p2, p1)
    var v20 int = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return -30
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 0
    var v3 int8 = v2
    var v4 uint32 = 2
    var v5 string = ("a" + "hello")
    var v6 uint32 = (^(v4))
    if (0 < 2) {
        switch 42 {
        case -1:
            var v21 string = v5
            _ = v21
        case 1:
            v22 := GenS3{F0: 728}
            _ = v22
        default:
            var v23 bool = false
            _ = v23
        }
        for fi24 := 0; fi24 < 16; fi24++ {
            v4 += (^(v6))
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 float64 = (((((-2.25 / 0.1) * ((1.5 / ((-2.25 * 0.1) / (1e10 / -2.25))) * 0.0)) + (1.5 - float64((3.14159 - 0.1)))) / 3.14159) / -2.25)
                _ = fi25
                _ = v26
            }
            v27 := []int{fi24}
            _ = fi24
            _ = v27
        }
        v28 := GenS3{F0: 0}
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v6 = v6
        var v30 float64 = ((0.1 / (1e10 + 3.14159)) + ((-2.25 / 1.5) - (0.0 / 1.0)))
        fmt.Println(v3, fi29)
        _ = fi29
        _ = v30
    }
    v5 = v5
    for fi31 := 0; fi31 < 17; fi31++ {
        v32 := GenS3{F0: 42}
        for fi33 := 0; fi33 < 19; fi33++ {
            if ((v6 << ((4294967295) & 63)) > (^(v6))) {
                var v34 float32 = (3.14159 + 0.1)
                v6 = 65535
                var v35 float64 = 0.1
                _ = v34
                _ = v35
            }
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    var v36 string = "\n\t"
    var v37 int32 = -251
    fmt.Println(v3, v37, v3)
    v38 := []int{51, 42, 32767}
    v36 = "hello"
    for fi39 := 0; fi39 < 1; fi39++ {
        v37 = (v37 + v37)
        _ = fi39
    }
    if true {
        var v40 uint32 = genFn0(29, 2)
        v6 += (v40 ^ genFn0(127, 1))
        _ = v40
    }
    v41 := GenS1{F0: 1, F1: 2}
    var v42 uint32 = 2
    var v43 float32 = 3.14159
    var v44 uint16 = 1
    fmt.Println(v4, v42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
