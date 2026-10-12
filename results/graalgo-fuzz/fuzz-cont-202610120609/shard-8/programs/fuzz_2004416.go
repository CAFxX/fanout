package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 int16) int64 {
    var v4 uint32 = 2
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(p0, p1, p0)
        v6 := GenS1{F0: 255, F1: (745 <= 2), F2: "foo bar"}
        _ = fi5
        _ = v6
    }
    fmt.Println(v4, p0, v4)
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 2, F1: true, F2: ("foo bar" + "hello")}
        var v9 int16 = (p1 >> ((297) & 63))
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        v10 := []int{128, -795, 2, -182}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v4
    return int64(v4)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    v3 := []string{"\n\t"}
    v2 += v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1[933] = 855
    var v11 uint16 = 900
    var v12 uint32 = 1
    var v13 uint8 = 42
    v3 = append(v3, "a")
    v14 := GenS2{F0: (723 != 725)}
    var v15 float32 = float32((1.5 - -2.25))
    if true {
        switch 127 {
        case -415:
            v16 := GenS3{F0: (v11 - (v11 % ((65535) | 1))), F1: "foo bar"}
            _ = v16
        }
        v2 = int16(v11)
    } else {
        v12 = v12
        v15 += 0.0
        v2 = (v2 / ((v2) | 1))
    }
    if (((!true) || true) || true) {
        v2 -= -1
        v13 -= 125
    } else {
        _ = len(v1)
        if (!true) {
            if (true && true) {
                var v17 uint32 = 679
                _ = v17
            } else {
                fmt.Println(v11, v13, v15)
            }
            switch 643 {
            case 32768:
                fmt.Println(v15, v15, v2)
            case 651:
                if false {
                    v11 += (v11 | v11)
                    v2 = (v2 / ((v2) | 1))
                } else {
                    v18 := GenS2{F0: ((v2 | v2) > -1)}
                    _ = v18
                }
            }
        }
        v2 = 1
    }
    if false {
        for fi19 := 0; fi19 < 6; fi19++ {
            _ = len(v3)
            _ = fi19
        }
    }
    v20 := GenS3{F0: v11, F1: ("a" + "\n\t")}
    var v21 bool = (true || true)
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS4{F0: (v13 * v13)}
        v12 += (v12 / (((^(v12))) | 1))
        _ = fi22
        _ = v23
    }
    _ = len(v1)
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
