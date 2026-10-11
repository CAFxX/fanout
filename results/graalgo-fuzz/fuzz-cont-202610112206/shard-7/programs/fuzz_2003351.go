package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 int32, p2 uint64) int16 {
    var v7 float32 = (1e10 * (1.5 - 0.1))
    fmt.Println(v7)
    v8 := []int{}
    v9 := GenS1{F0: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint) uint32 {
    v10 := []int{32767, 32767, -285}
    var v11 uint16 = 32767
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v10) > 0 {
            v10[0] = 4294967295
        }
        _ = fi12
    }
    var v13 int = 2147483647
    _ = p0
    _ = v10
    _ = v11
    _ = v13
    return 379
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := []string{"\n\t", "", "foo bar"}
    v5 := make(map[int]int)
    var v6 uint16 = 1
    for fi14 := 0; fi14 < 19; fi14++ {
        v6 += (v6 & 255)
        v15 := GenS1{F0: (^(v6))}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v6, fi16)
        _ = fi16
    }
    v17 := GenS1{F0: (v6 % (((^(v6))) | 1))}
    v6 -= v6
    v6 = (v6 | 474)
    var v18 float64 = (-2.25 / 0.1)
    var v19 uint32 = 2147483647
    v18 += 1.5
    _ = len(v4)
    v5[128] = 9223372036854775807
    var v20 int64 = -1
    fmt.Println(v20, v20, v20)
    v21 := GenS2{F0: 0}
    v18 = v18
    switch ("" + (("" + (("foo bar" + "hello") + "")) + "")) {
    case "\n\t":
        v20 = v20
    case "a":
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := v3[int(v19)]
            v4 = append(v4, ("\n\t" + "foo bar"))
            _ = fi22
            _ = v23
        }
    case "foo bar":
        v24 := GenS2{F0: 24}
        _ = v24
    }
    v19 += (v19 * v19)
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v6, v19)
        if (v18 < v18) {
            v4 = append(v4, "hello")
            v26 := GenS3{F0: genFn0(v6, -693, 18446744073709551615)}
            var v27 uint8 = 127
            _ = v26
            _ = v27
        } else {
            switch (("foo bar" + "") + "\n\t") {
            case "foo bar":
                v3[4294967295] = fi25
            }
        }
        _ = fi25
    }
    delete(v3, 352)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
