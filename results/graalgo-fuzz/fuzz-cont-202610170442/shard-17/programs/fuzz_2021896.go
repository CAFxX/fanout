package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{878}
    var v3 bool = true
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int8 = 127
        v5 -= v5
        v5 += v5
        _ = fi4
        _ = v5
    }
    v3 = (!(!((0.0 - (3.14159 / 1.0)) >= 0.0)))
    v3 = false
    v3 = v3
    v3 = (83 <= 0)
    v6 := GenS1{F0: 42}
    var v7 uint8 = 2
    fmt.Println(v7)
    v7 = v7
    var v8 float64 = (1e10 * float64(1e10))
    for fi9 := 0; fi9 < 13; fi9++ {
        v8 += v8
        _ = fi9
    }
    v10 := v1[-2147483648]
    switch 2147483647 {
    case -2147483648:
        v7 = 0
    case 2147483647:
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(v8, v3, fi11)
            _ = len(v2)
            _ = fi11
        }
    default:
        fmt.Println(v8)
    }
    if v10 {
        v7 -= (v7 - (v7 << ((255) & 63)))
        var v12 uint = 126
        _ = v12
    } else {
        v13 := v1[-9223372036854775808]
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(fi14)
        var v15 int64 = 1
        if ((v8 - v8) <= -2.25) {
            if v3 {
                var v16 uint = 42
                for fi17 := 0; fi17 < 1; fi17++ {
                    v3 = v3
                    switch (fi17 | (-(fi17))) {
                    case -327:
                        var v18 uint8 = (v7 << ((65535) & 63))
                        _ = v18
                    }
                    var v19 int64 = -745
                    _ = fi17
                    _ = v19
                }
                _ = v16
            }
            var v20 int16 = 1
            switch (("a" + "") + ("a" + "foo bar")) {
            case "":
                if true {
                    v15 -= -967
                }
            case "\n\t":
                delete(v1, -9223372036854775808)
            case "a":
                if v10 {
                    fmt.Println(v10, fi14)
                }
            }
            _ = v20
        } else {
            v3 = (!true)
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v10)
    for fi21 := 0; fi21 < 10; fi21++ {
        switch (fi21 - 128) {
        case 434:
            for fi22 := 0; fi22 < 14; fi22++ {
                switch -1 {
                case 965:
                    v23 := GenS2{F0: 2, F1: (1 <= -9223372036854775808), F2: 0}
                    _ = v23
                case 127:
                    var v24 int16 = -32768
                    _ = v24
                case -2147483648:
                    v7 = (v7 & (v7 * v7))
                }
                v3 = v10
                _ = fi22
            }
        case 170:
            v10 = v10
        case 2147483647:
            fmt.Println(v8)
        }
        var v25 bool = v3
        if false {
            for fi26 := 0; fi26 < 18; fi26++ {
                if (!v25) {
                    v3 = v25
                    _ = len(v2)
                    delete(v1, fi21)
                }
                v10 = (!((-(fi26)) < (fi26 * -190)))
                v8 = float64(3.14159)
                _ = fi26
            }
        } else {
            v8 += v8
        }
        _ = fi21
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
