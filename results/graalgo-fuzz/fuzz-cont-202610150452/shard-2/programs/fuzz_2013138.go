package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
func genFn0(p0 int16) uint8 {
    if true {
        v5 := []int{-187}
        _ = v5
    } else {
        if true {
            var v6 uint = 18446744073709551615
            var v7 uint16 = 278
            _ = v6
            _ = v7
        } else {
            var v8 string = ("\n\t" + "\n\t")
            v8 = (v8 + (v8 + (v8 + "foo bar")))
            _ = v8
        }
        var v9 int16 = -637
        if false {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 string = "hello"
                v12 := []int{}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 float32 = 1e10
            v14 := []int{128, -772, 1}
            _ = v13
            _ = v14
        }
        _ = v9
    }
    switch "" {
    case "foo bar":
        v15 := []int{32768, 368}
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int64 = 0
            _ = fi16
            _ = v17
        }
    case "\n\t":
        var v18 string = ""
        _ = v18
    }
    if (!(9223372036854775807 <= -632)) {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int8 = 0
            var v21 string = ""
            v20 = v20
            _ = fi19
            _ = v20
            _ = v21
        }
        if ((4294967295 <= 891) || false) {
            var v22 string = ("foo bar" + ("" + "\n\t"))
            p0 -= p0
            p0 += 1000
            _ = v22
        }
        fmt.Println(p0)
    }
    v23 := GenS1{F0: 2147483647, F1: 65535, F2: 1}
    _ = p0
    _ = v23
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{171, -1}
    var v3 uint64 = 310
    var v4 int64 = -555
    _ = len(v2)
    for fi24 := 0; fi24 < 20; fi24++ {
        switch "foo bar" {
        case "":
            v4 = 0
        }
        _ = fi24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v4 -= (-(v4))
        var v26 int8 = 0
        var v27 float32 = (1.5 - 1e10)
        _ = fi25
        _ = v26
        _ = v27
    }
    v3 += v3
    v3 += v3
    var v28 uint8 = uint8(v3)
    var v29 int8 = 108
    v3 = (v3 - (^(v3)))
    var v30 float32 = 0.0
    switch 127 {
    case 32767:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 += v3
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        v28 += 127
        _ = fi31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        delete(v1, 0)
        v3 -= v3
        if (v29 < v29) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v33 uint8 = 127
            _ = v33
        }
        _ = fi32
    }
    var v34 uint8 = (v28 << ((255) & 63))
    v34 = v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
}
