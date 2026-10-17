package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 bool, p2 uint8) uint64 {
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(p2)
        _ = fi8
    }
    v9 := GenS1{F0: 437, F1: "hello", F2: "a"}
    p1 = false
    for fi10 := 0; fi10 < 4; fi10++ {
        switch "foo bar" {
        case "a":
            p1 = false
        case "hello":
            v11 := []int{32768}
            _ = v11
        case "\n\t":
            v12 := []int{(fi10 - fi10)}
            _ = v12
        }
        if ((fi10 * (fi10 * (fi10 ^ (fi10 << ((18446744073709551615) & 63))))) > 32768) {
            p1 = (!false)
            v13 := GenS1{F0: 0, F1: "", F2: "a"}
            _ = v13
        }
        switch 32767 {
        case -2147483648:
            fmt.Println(v7)
        case 42:
            if false {
                v14 := GenS1{F0: 0, F1: "a", F2: ("\n\t" + (("a" + (("foo bar" + "") + "")) + "\n\t"))}
                var v15 int64 = 0
                p1 = ("\n\t" != "\n\t")
                _ = v14
                _ = v15
            }
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 2
}

func main() {
    var v1 uint32 = 53
    var v2 uint64 = 2
    var v3 float32 = (-2.25 + ((1.5 + ((-2.25 + 0.0) + 3.14159)) / (-2.25 - (float32((-2.25 - 1e10)) - (1e10 + ((0.0 / (1e10 / 1.5)) * 0.0))))))
    var v4 int = -2147483648
    v5 := []uint32{890, (v1 | v1), 2}
    v6 := make(map[uint32]int)
    v16 := v6[(v1 % ((v1) | 1))]
    switch (v4 ^ (v16 / (((^(v16))) | 1))) {
    case 4294967295:
        var v17 int16 = 32767
        _ = v17
    case 9223372036854775807:
        if (v2 == v2) {
            var v18 uint = 255
            fmt.Println(v4, v2, v18)
            v4 = (v4 + (v4 ^ (v16 / ((127) | 1))))
            _ = v18
        }
    case 1:
        switch "hello" {
        case "":
            v1 = (v1 | 1)
        case "foo bar":
            fmt.Println(v4)
        case "hello":
            v1 = 4294967295
        }
    }
    v19 := GenS2{F0: 0}
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS1{F0: 32767, F1: "a", F2: "\n\t"}
        for fi22 := 0; fi22 < 6; fi22++ {
            fmt.Println(v16, v4, fi20)
            fmt.Println(v1, fi20, fi22)
            _ = fi22
        }
        v16 = (v4 % (((-(v4))) | 1))
        _ = fi20
        _ = v21
    }
    v6[(v1 - v1)] = (v16 * (^(v16)))
    v23 := v6[(v1 / (((v1 - (v1 - (v1 | (v1 >> ((1) & 63)))))) | 1))]
    var v24 string = ("hello" + "a")
    v25 := v6[428]
    v23 += (v16 % (((v16 - -9223372036854775808)) | 1))
    v3 += float32(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
