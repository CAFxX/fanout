package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 2147483647
    switch 9223372036854775807 {
    case 1:
        v12 := GenS1{F0: 28}
        _ = v12
    case -2147483648:
        v13 := []int{1}
        _ = v13
    default:
        switch -2147483648 {
        case 32767:
            var v14 string = "a"
            _ = v14
        }
    }
    if false {
        v15 := []int{1, 210, -873, 345}
        _ = v15
    } else {
        var v16 float32 = 1.5
        v17 := []int{120}
        _ = v16
        _ = v17
    }
    _ = v11
    return -669
}
func genFn1() int8 {
    v18 := []int{32767}
    var v19 bool = (false || ((true && true) || true))
    fmt.Println(v19)
    _ = v18
    _ = v19
    return -96
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("" + "a")
    v3 := []string{(v2 + ""), v2, v2}
    var v4 int8 = 1
    var v5 int64 = 554
    v6 := []string{v2, v2, v2, v2}
    if false {
        v20 := GenS1{F0: v4}
        _ = v20
    } else {
        v4 += v4
    }
    switch v2 {
    case "hello":
        if false {
            var v21 float32 = (-2.25 / 1.0)
            var v22 int = -249
            _ = v21
            _ = v22
        } else {
            _ = len(v1)
            if false {
                var v23 string = v2
                var v24 int64 = (-(v5))
                for fi25 := 0; fi25 < 16; fi25++ {
                    fmt.Println(v4)
                    v24 = (v5 << ((274) & 63))
                    v26 := GenS1{F0: (v4 / ((v4) | 1))}
                    _ = fi25
                    _ = v26
                }
                _ = v23
                _ = v24
            }
        }
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 float32 = 1.5
            if true {
                v4 += 26
            } else {
                var v30 float32 = ((v29 - (v29 + v29)) * (1.5 / v29))
                fmt.Println(v30)
                _ = v30
            }
            if (255 == 1) {
                var v31 bool = (!true)
                var v32 uint32 = 1
                v33 := GenS1{F0: v4}
                v4 += v4
                v34 := GenS2{F0: ("a" != ("" + (v2 + "")))}
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi28
            _ = v29
        }
        switch "\n\t" {
        case "\n\t":
            var v35 int32 = -119
            _ = v35
        }
        v2 = ("foo bar" + ("foo bar" + v2))
        _ = fi27
    }
    fmt.Println(v2, v5)
    v4 += (v4 & (v4 ^ (-(v4))))
    v4 = genFn1()
    switch (v2 + v2) {
    case "hello":
        v36 := v1[-210]
        _ = v36
    case "foo bar":
        v5 -= v5
    }
    v2 = v2
    v5 = v5
    v37 := GenS1{F0: -91}
    switch ("foo bar" + v2) {
    case "foo bar":
        v38 := GenS2{F0: true}
        _ = v38
    case "hello":
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 float32 = -2.25
            fmt.Println(v2)
            _ = fi39
            _ = v40
        }
    case "a":
        v4 = (v4 >> ((1) & 63))
    default:
        for fi41 := 0; fi41 < 18; fi41++ {
            if (int32(fi41) >= genFn0()) {
                switch (v2 + "\n\t") {
                case "hello":
                    v4 -= v4
                case "":
                    fmt.Println(v4, fi41, v2)
                }
            } else {
                var v42 int64 = (v5 * 1)
                delete(v1, (-(fi41)))
                v5 = 0
                _ = v42
            }
            v43 := GenS3{F0: false, F1: 0, F2: 158}
            _ = fi41
            _ = v43
        }
    }
    switch v2 {
    case "hello":
        var v44 int64 = v5
        _ = v44
    default:
        fmt.Println(v4, v2)
    }
    if true {
        v2 = (v2 + (v2 + ("a" + v2)))
    } else {
        var v45 int32 = 1
        fmt.Println(v2, v5)
        _ = v45
    }
    var v46 float64 = 1e10
    var v47 uint32 = 505
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
