package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int32; F1 uint16; F2 uint32 }

func main() {
    v1 := []uint32{0, 65535}
    var v2 float64 = 3.14159
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v2 += (float64(3.14159) * (0.1 - (1e10 - 0.1)))
    v4[0] = "a"
    v5 := GenS1{F0: 42, F1: 1}
    v2 = 0.1
    v6 := GenS2{F0: true, F1: ""}
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS2{F0: true, F1: "hello"}
        for fi9 := 0; fi9 < 20; fi9++ {
            if ((-2.25 != float32((1e10 - 1e10))) || true) {
                v2 = v2
                fmt.Println(fi7)
            } else {
                var v10 uint8 = 223
                _ = v10
            }
            v11 := GenS2{F0: (!(-121 < -249)), F1: ""}
            _ = fi9
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if false {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
        v13 := GenS3{F0: -1, F1: 42, F2: 2}
        v2 = 1.5
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "a":
            v2 -= 1.5
        default:
            v2 = 0.1
        }
        fmt.Println(v2)
        v15 := GenS2{F0: true, F1: ("a" + "hello")}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        delete(v4, 65535)
        _ = fi16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v2 = 1.5
        _ = fi17
    }
    v18 := GenS1{F0: 255, F1: 588}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
