package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint16) int32 {
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 += (^(p0))
        var v6 int32 = 2147483647
        _ = fi5
        _ = v6
    }
    switch ("" + "") {
    case "a":
        for fi7 := 0; fi7 < 7; fi7++ {
            p0 -= p0
            p1 = 2
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int16 = -273
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "":
        for fi10 := 0; fi10 < 15; fi10++ {
            p0 += p0
            v11 := []int{(fi10 << ((282) & 63))}
            p0 -= p0
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    return -16
}
func genFn1(p0 uint64, p1 bool, p2 int8) uint64 {
    for fi12 := 0; fi12 < 12; fi12++ {
        if true {
            v13 := []int{42, 32767, (fi12 >> ((227) & 63))}
            v14 := GenS1{F0: p0}
            var v15 uint16 = uint16(p0)
            _ = v13
            _ = v14
            _ = v15
        } else {
            p0 = 4294967295
            v16 := []int{fi12, fi12, fi12, (fi12 % ((fi12) | 1))}
            _ = v16
        }
        v17 := GenS1{F0: (^(p0))}
        v18 := GenS1{F0: 4294967295}
        _ = fi12
        _ = v17
        _ = v18
    }
    var v19 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return (p0 & p0)
}

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]uint32)
    var v3 bool = (3.14159 <= -2.25)
    var v4 int64 = 42
    switch -1 {
    case 4294967295:
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int8 = -128
            _ = fi20
            _ = v21
        }
    case 127:
        if ((1 > 397) && v3) {
            v3 = true
        } else {
            var v22 uint32 = 4294967295
            _ = v22
        }
    }
    v4 -= v4
    if v3 {
        if len(v1) > 0 {
            v1[0] = ("a" + "hello")
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 int16 = -32768
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        v4 += (v4 + v4)
    } else {
        switch "foo bar" {
        case "hello":
            var v26 string = "\n\t"
            _ = v26
        case "":
            v4 = v4
        }
        var v27 uint32 = 65535
        _ = v27
    }
    var v28 uint8 = 1
    var v29 float32 = 1e10
    fmt.Println(v3)
    v28 -= 133
    switch "" {
    case "hello":
        if (((!false) || true) || v3) {
            v4 += (v4 >> ((255) & 63))
        } else {
            if false {
                v30 := GenS1{F0: 369}
                for fi31 := 0; fi31 < 15; fi31++ {
                    v1 = append(v1, ("a" + ("\n\t" + "hello")))
                    var v32 int16 = 431
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
            var v33 uint16 = 2
            for fi34 := 0; fi34 < 4; fi34++ {
                v28 = v28
                _ = fi34
            }
            _ = v33
        }
    case "a":
        v35 := GenS1{F0: 1}
        _ = v35
    default:
        fmt.Println(v28)
    }
    switch "hello" {
    case "\n\t":
        v3 = (2 == 2147483647)
    case "":
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    v2[0] = 686
    var v36 string = "hello"
    var v37 int32 = int32(v28)
    if v3 {
        var v38 int64 = 9223372036854775807
        var v39 int = 698
        _ = v38
        _ = v39
    }
    fmt.Println(v37, v3)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
