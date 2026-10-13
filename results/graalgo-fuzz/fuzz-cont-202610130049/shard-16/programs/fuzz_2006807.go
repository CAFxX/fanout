package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64, p2 uint32) uint {
    if p0 {
        v6 := GenS1{F0: "foo bar", F1: 48}
        _ = v6
    } else {
        if p0 {
            var v7 int64 = -1
            var v8 uint8 = 139
            _ = v7
            _ = v8
        }
        p0 = (p0 || (!p0))
    }
    var v9 uint64 = (p1 ^ p1)
    if (127 < -2147483648) {
        p0 = (p0 || p0)
    }
    var v10 uint8 = 4
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return uint(p2)
}
func genFn1(p0 uint32, p1 uint16, p2 uint32) uint16 {
    if false {
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{(fi12 % (((fi12 & (fi11 + fi12))) | 1))}
                _ = fi12
                _ = v13
            }
            var v14 bool = true
            _ = fi11
            _ = v14
        }
        var v15 float64 = 0.1
        _ = v15
    } else {
        p2 = 733
        p0 = p2
        p2 = 0
    }
    v16 := []int{127, -483}
    var v17 uint64 = 42
    fmt.Println(v17)
    v18 := GenS2{F0: 310, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return 255
}

func main() {
    var v1 uint16 = 113
    var v2 int64 = 1
    v3 := []string{}
    v4 := make(map[uint32]bool)
    var v5 uint16 = 2
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 bool = (!((((false && (!("a" < ("hello" + ("\n\t" + ("foo bar" + "")))))) || true) || (v1 > (v1 * (v5 >> ((0) & 63))))) || (!(0.1 <= ((-2.25 - (((float32(1.5) / -2.25) + (1.5 / 0.1)) - (0.1 / (-2.25 * ((0.0 + (1e10 * 3.14159)) * (1.5 + 0.0)))))) + 1.5)))))
        v21 := GenS3{F0: 4294967295}
        _ = fi19
        _ = v20
        _ = v21
    }
    if (((("hello" + "\n\t") + "") + "a") != "\n\t") {
        var v22 uint = 18446744073709551615
        v23 := GenS1{F0: "", F1: (v2 % ((v2) | 1))}
        if false {
            v24 := GenS2{F0: v22, F1: 42}
            _ = v24
        }
        _ = v22
        _ = v23
    } else {
        v25 := GenS1{F0: "hello", F1: (-(v2))}
        _ = v25
    }
    v26 := GenS3{F0: 812}
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            fmt.Println(fi28, v5, v5)
            var v29 float32 = float32(-2.25)
            var v30 uint64 = 183
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = len(v3)
        v3 = append(v3, "hello")
        _ = fi27
    }
    v31 := GenS1{F0: ("foo bar" + "a"), F1: (v2 << ((1) & 63))}
    var v32 uint16 = v5
    v1 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!((!true) && false)) {
        v1 = 1
        var v33 uint8 = 127
        switch "" {
        case "a":
            v34 := v4[2]
            _ = v34
        case "hello":
            var v35 int8 = -1
            _ = v35
        case "":
            v32 = v5
        default:
            v2 += v2
        }
        _ = v33
    }
    var v36 bool = true
    fmt.Println(v5, v2, v5)
    var v37 int = -9223372036854775808
    var v38 float64 = (1.5 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
