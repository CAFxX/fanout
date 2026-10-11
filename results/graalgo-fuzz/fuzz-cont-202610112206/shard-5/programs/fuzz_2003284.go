package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16) int16 {
    p0 -= (p0 % ((p0) | 1))
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = 895
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 0}
        fmt.Println(p0, p1)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 int32, p1 uint8) uint64 {
    v7 := GenS1{F0: 175}
    var v8 int8 = -1
    if (((!true) && (1 == 927)) && (true || true)) {
        v8 = int8(p0)
        v9 := []int{}
        if false {
            v10 := GenS1{F0: 2147483647}
            _ = v10
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    switch -9223372036854775808 {
    case 1:
        var v11 int16 = v1
        _ = v11
    }
    if true {
        if ((!((v1 / (((-(v1))) | 1)) <= (-(v1)))) && false) {
            v1 += 1
        }
        if (!((-2.25 > 0.0) && (100 <= 0))) {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := v3[(fi12 | 32767)]
                switch 1 {
                case -360:
                    var v14 uint8 = 2
                    _ = v14
                case -9223372036854775808:
                    var v15 uint32 = 4294967295
                    _ = v15
                default:
                    var v16 int32 = -891
                    _ = v16
                }
                var v17 int16 = -1
                _ = fi12
                _ = v13
                _ = v17
            }
        }
        if false {
            delete(v2, (("foo bar" + ("\n\t" + "\n\t")) + ((("foo bar" + "a") + "a") + ("" + ("a" + "foo bar")))))
            v1 += v1
        }
    } else {
        var v18 int16 = v1
        var v19 uint16 = 32767
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int = (-(fi20))
        _ = fi20
        _ = v21
    }
    v1 = (v1 << ((121) & 63))
    if (false || true) {
        switch "\n\t" {
        case "\n\t":
            var v22 uint64 = genFn1(-137, 0)
            _ = v22
        case "":
            fmt.Println(v1)
        case "hello":
            v1 -= v1
        }
        var v23 uint = 2
        _ = v23
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            if (false || (("foo bar" >= ("" + ("hello" + ("" + (("" + "\n\t") + "\n\t"))))) || (0 <= 0))) {
                v25 := GenS1{F0: 2}
                var v26 string = (("foo bar" + "a") + "hello")
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 11; fi27++ {
                    fmt.Println(v1, v1)
                    fmt.Println(fi24, v1, fi27)
                    v28 := []int{}
                    _ = fi27
                    _ = v28
                }
                v29 := []int{-124, (^(fi24))}
                _ = v29
            }
            _ = fi24
        }
        var v30 float32 = (1e10 + 0.1)
        var v31 int = 32768
        _ = v30
        _ = v31
    }
    var v32 int16 = v1
    if ((!(!(true || false))) && false) {
        var v33 uint = 18446744073709551615
        v34 := GenS1{F0: 638}
        v35 := GenS2{F0: 0}
        _ = v33
        _ = v34
        _ = v35
    }
    v1 += (v32 << ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
}
