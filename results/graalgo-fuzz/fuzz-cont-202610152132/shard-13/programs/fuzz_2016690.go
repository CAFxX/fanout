package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16) uint32 {
    v7 := GenS1{F0: (p0 >> ((1) & 63))}
    p0 -= p0
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := []int{fi8, fi9}
                _ = fi9
                _ = v10
            }
        }
        _ = fi8
    }
    var v11 uint32 = 2
    if false {
        if true {
            v11 = (v11 - 2)
            v12 := []int{44}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                fmt.Println(fi13, v11)
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := GenS1{F0: p0}
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            if false {
                var v16 float32 = (((float32((0.0 - 0.1)) * 1e10) * 1.5) - 0.1)
                _ = v16
            } else {
                p0 = (p0 ^ 2)
            }
        }
    }
    _ = p0
    _ = v11
    _ = v7
    return (v11 + (v11 | (v11 * (v11 | v11))))
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    var v4 int = 1
    var v5 string = "a"
    var v6 uint64 = 0
    var v17 int16 = 1
    for fi18 := 0; fi18 < 18; fi18++ {
        switch (("foo bar" + "hello") + v5) {
        case "hello":
            if (!((v6 + (v6 * v6)) <= (v6 & (v6 % ((v6) | 1))))) {
                v19 := GenS1{F0: 2}
                _ = v19
            }
        default:
            var v20 int64 = 42
            _ = v20
        }
        switch fi18 {
        case 0:
            _ = len(v3)
        case 32767:
            v3[1] = (v1 * (v1 ^ (v1 - v1)))
        case 9223372036854775807:
            var v21 uint32 = 4294967295
            _ = v21
        default:
            for fi22 := 0; fi22 < 10; fi22++ {
                fmt.Println(v5)
                var v23 string = ("" + "foo bar")
                fmt.Println(fi18)
                _ = fi22
                _ = v23
            }
        }
        _ = fi18
    }
    v24 := GenS1{F0: uint16(v1)}
    for fi25 := 0; fi25 < 5; fi25++ {
        switch (v5 + v5) {
        case "hello":
            v26 := []int{v4}
            _ = v26
        default:
            v27 := v3[genFn0(718)]
            _ = v27
        }
        v28 := GenS1{F0: 181}
        v29 := []int{0}
        _ = fi25
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        if false {
            fmt.Println(v5, v17, v6)
            v31 := []int{42}
            _ = v31
        }
        switch v5 {
        case "hello":
            v32 := GenS2{F0: genFn0(369)}
            _ = v32
        }
        v5 = v2
        _ = fi30
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v17 = v17
        _ = fi33
    }
    v5 = v2
    v4 += v4
    var v34 uint8 = 255
    v5 = v2
    switch v4 {
    case 0:
        var v35 int16 = v17
        _ = v35
    default:
        fmt.Println(v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
