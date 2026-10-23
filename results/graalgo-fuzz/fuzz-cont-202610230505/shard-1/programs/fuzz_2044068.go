package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0() uint {
    var v5 string = ((("foo bar" + ("" + "hello")) + "foo bar") + "\n\t")
    var v6 float32 = (1.5 + ((1.5 / 1e10) * ((0.1 - 1e10) - (1.5 + 1e10))))
    for fi7 := 0; fi7 < 2; fi7++ {
        v5 = "\n\t"
        for fi8 := 0; fi8 < 19; fi8++ {
            switch ("a" + (v5 + "a")) {
            case "hello":
                fmt.Println(v5)
            }
            v6 = 0.1
            var v9 uint64 = 42
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 156
}
func genFn1(p0 uint16) uint16 {
    p0 += (p0 & (p0 % ((p0) | 1)))
    v10 := GenS1{F0: "a", F1: 1}
    p0 = (p0 >> ((507) & 63))
    fmt.Println(p0)
    if true {
        p0 = p0
        var v11 int = 706
        if (true && (!((float64(((float32((-2.25 * (-2.25 * 1.5))) * 0.0) + float32((1.5 / (0.0 - 1e10))))) / float64((3.14159 - float32((3.14159 * (0.0 / ((0.0 - -2.25) - float64(0.0)))))))) < (0.0 + (-2.25 * float64(1.5)))))) {
            v11 = v11
            var v12 uint16 = (p0 + (p0 ^ (p0 / ((p0) | 1))))
            var v13 int16 = 0
            _ = v12
            _ = v13
        }
        _ = v11
    } else {
        if (!(p0 == p0)) {
            var v14 uint32 = 2
            v14 = 2147483647
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            if (113 != 918) {
                p0 -= (p0 << ((770) & 63))
                p0 += (p0 << ((0) & 63))
                p0 = p0
            } else {
                v16 := GenS1{F0: (("\n\t" + ((("\n\t" + "hello") + ("\n\t" + ("a" + "hello"))) + ("" + (("a" + "foo bar") + "\n\t")))) + "foo bar"), F1: 42}
                _ = v16
            }
            p0 -= 722
            fmt.Println(fi15, fi15)
            _ = fi15
        }
        p0 -= (p0 << ((65535) & 63))
    }
    _ = p0
    _ = v10
    return p0
}

func main() {
    v1 := []int{32767, 4294967295, -1, -314}
    var v2 uint = 0
    var v3 uint16 = 781
    var v4 uint = v2
    var v17 uint16 = 0
    v18 := GenS2{F0: 403}
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 float32 = 0.1
        _ = fi19
        _ = v20
    }
    v1 = append(v1, 128)
    var v21 uint16 = uint16(v4)
    var v22 float64 = (3.14159 / 1e10)
    v23 := GenS3{F0: 2, F1: 127}
    v22 += 0.1
    switch 127 {
    case 42:
        for fi24 := 0; fi24 < 7; fi24++ {
            v17 -= (v3 / ((v21) | 1))
            _ = len(v1)
            var v25 uint16 = v21
            _ = fi24
            _ = v25
        }
    default:
        v2 = v4
    }
    v22 = (v22 / float64((((0.1 - 0.0) / 1.5) / 1.5)))
    var v26 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
