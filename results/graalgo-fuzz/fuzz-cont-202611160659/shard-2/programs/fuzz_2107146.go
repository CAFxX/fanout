package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int8) uint8 {
    var v6 float64 = float64(0.1)
    v7 := []int{p0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]bool)
    var v3 int16 = 537
    v4 := []uint32{2147483647, 867}
    v5 := []string{"\n\t"}
    v8 := GenS1{F0: v3}
    v9 := GenS1{F0: v3}
    if true {
        if (int8(v1) != -43) {
            v3 = 576
            v3 = (v3 % (((^(v3))) | 1))
        }
        v10 := GenS2{F0: 255}
        if ((true || (-1 <= -628)) && (0 != 128)) {
            var v11 int32 = -1
            switch "a" {
            case "a":
                fmt.Println(v11, v11, v1)
            }
            _ = v11
        }
        _ = v10
    }
    switch "\n\t" {
    case "\n\t":
        v1 = v1
    }
    v2[2] = true
    v12 := GenS1{F0: 32767}
    _ = len(v2)
    switch 42 {
    case 32767:
        v1 -= (v1 >> ((926) & 63))
    }
    v1 += 42
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
