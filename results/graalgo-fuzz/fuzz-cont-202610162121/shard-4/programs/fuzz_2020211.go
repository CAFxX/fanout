package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 uint32) uint16 {
    p1 -= p1
    for fi5 := 0; fi5 < 16; fi5++ {
        switch "foo bar" {
        case "\n\t":
            var v6 int32 = 0
            _ = v6
        }
        switch (("" + "") + "") {
        case "\n\t":
            p0 = (p0 - -9223372036854775808)
        case "foo bar":
            p1 = (p1 << ((0) & 63))
        case "hello":
            v7 := GenS1{F0: p0}
            _ = v7
        }
        p1 = p1
        _ = fi5
    }
    p0 += (p0 + (p0 / ((-729) | 1)))
    p0 += p0
    p1 += (p1 ^ p1)
    _ = p0
    _ = p1
    return 604
}
func genFn1() uint32 {
    var v8 string = "hello"
    v9 := []int{}
    v8 = v8
    v8 = (v8 + "hello")
    _ = v8
    _ = v9
    return 119
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]int)
    var v3 int = 128
    var v4 int = (-(v3))
    fmt.Println(v1, v1, v4)
    v3 = 464
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(fi11, v3, fi10)
            v12 := v2[("foo bar" + "foo bar")]
            _ = fi11
            _ = v12
        }
        if false {
            var v13 bool = (fi10 < (-(fi10)))
            _ = v13
        } else {
            v4 = (v3 << ((265) & 63))
            v4 = (^(v3))
        }
        var v14 uint8 = 1
        _ = fi10
        _ = v14
    }
    var v15 uint = 0
    var v16 uint64 = 309
    v17 := []int{(v4 % (((v3 + -48)) | 1)), 127, v3}
    fmt.Println(v3, v16, v4)
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int = (v3 - (fi18 + (v4 + 4294967295)))
        switch "a" {
        case "\n\t":
            v20 := GenS1{F0: 1}
            _ = v20
        case "hello":
            v21 := GenS2{F0: (v15 + 255)}
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v1 = (v1 - (v1 - 0.0))
    switch -9223372036854775808 {
    case 42:
        switch (v4 + (v3 % ((-1) | 1))) {
        case 42:
            if true {
                if (-681 >= 1) {
                    v2["a"] = -1
                }
            }
        case 9223372036854775807:
            for fi22 := 0; fi22 < 17; fi22++ {
                v1 += (1e10 - (v1 * (float32(0.1) / float32(1e10))))
                _ = fi22
            }
        }
    case 127:
        v3 = (v3 * (^(v3)))
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        _ = len(v2)
        v24 := GenS2{F0: (v15 ^ 255)}
        _ = fi23
        _ = v24
    }
    v16 = (v16 - v16)
    v3 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
