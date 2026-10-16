package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8) int8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        p0 = p0
        fmt.Println(p0, p0)
        switch "hello" {
        case "hello":
            fmt.Println(p0)
        }
        _ = fi4
    }
    if false {
        fmt.Println(p0)
        v5 := []int{32768, -954, -12}
        if false {
            switch "a" {
            case "foo bar":
                p0 += (^(p0))
            }
        } else {
            fmt.Println(p0)
            var v6 bool = true
            _ = v6
        }
        _ = v5
    } else {
        switch 113 {
        case 32767:
            p0 += 79
        }
        if (true || (42 != 1)) {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int64 = 1
                if false {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        v8 = 998
                        _ = fi9
                    }
                }
                p0 += p0
                _ = fi7
                _ = v8
            }
            v10 := []int{-386, -486, -190}
            v11 := GenS1{F0: true, F1: -1}
            _ = v10
            _ = v11
        } else {
            fmt.Println(p0)
        }
    }
    _ = p0
    return 127
}

func main() {
    var v1 int = -402
    var v2 int64 = -500
    v3 := []uint32{}
    v2 -= (v2 >> ((4294967295) & 63))
    switch v1 {
    case -202:
        switch "\n\t" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 21
            }
        }
    case 2147483647:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v12 := GenS2{F0: 0}
    switch (("hello" + (("\n\t" + "\n\t") + (("" + ("\n\t" + "a")) + "foo bar"))) + "hello") {
    case "\n\t":
        if (!(1 >= 226)) {
            fmt.Println(v2, v2)
        }
    case "":
        var v13 uint8 = 128
        _ = v13
    default:
        var v14 uint = 0
        _ = v14
    }
    v2 -= (v2 * v2)
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(fi15)
            v17 := GenS1{F0: false, F1: -9223372036854775808}
            switch (v1 * (v1 % ((v1) | 1))) {
            case 0:
                var v18 uint8 = 127
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = len(v3)
        _ = fi15
    }
    v1 += (v1 + v1)
    if true {
        var v19 int8 = 127
        v19 = v19
        v20 := GenS2{F0: 65535}
        _ = v19
        _ = v20
    }
    var v21 uint32 = 261
    _ = len(v3)
    v22 := GenS3{F0: (v21 + (v21 % ((v21) | 1)))}
    if ("a" <= ("\n\t" + "")) {
        v1 += v1
    } else {
        var v23 float32 = 0.0
        v24 := GenS3{F0: v21}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        fmt.Println(v21)
        v26 := GenS3{F0: 2147483648}
        _ = fi25
        _ = v26
    }
    if false {
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(v21, v21)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi27
        }
        var v28 float64 = 1e10
        _ = v28
    } else {
        v1 += int(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
