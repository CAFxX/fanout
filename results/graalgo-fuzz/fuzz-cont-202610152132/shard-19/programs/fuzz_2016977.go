package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 int32 = 42
    var v3 uint = 288
    v4 := []string{"a", "\n\t"}
    var v5 uint8 = 2
    var v6 uint64 = v1
    fmt.Println(v2, v3, v3)
    v6 += v1
    if (100 <= -128) {
        v7 := GenS1{F0: (v5 & (v5 & v5))}
        _ = v7
    } else {
        v3 -= v3
        var v8 uint = v3
        v2 = v2
        _ = v8
    }
    switch (("hello" + "hello") + "a") {
    case "a":
        for fi9 := 0; fi9 < 1; fi9++ {
            switch (fi9 / (((fi9 + (fi9 % ((4294967295) | 1)))) | 1)) {
            case 205:
                v2 += (v2 >> ((309) & 63))
            case 0:
                v1 -= v6
            }
            _ = fi9
        }
    case "foo bar":
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint = v3
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 int8 = 98
                v2 -= (^(v2))
                _ = fi12
                _ = v13
            }
            var v14 bool = true
            _ = fi10
            _ = v11
            _ = v14
        }
    default:
        var v15 int32 = -501
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 -= uint(v5)
        fmt.Println(v3)
        var v17 string = ("a" + ((("" + "\n\t") + ("hello" + ("foo bar" + ("\n\t" + "foo bar")))) + ("hello" + ("" + ("" + "\n\t")))))
        _ = fi16
        _ = v17
    }
    if (!true) {
        v18 := GenS1{F0: 1}
        _ = v18
    } else {
        var v19 int32 = (v2 >> ((v3) & 63))
        for fi20 := 0; fi20 < 19; fi20++ {
            v4 = append(v4, "foo bar")
            _ = fi20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint32 = 4294967295
        v2 += (-(v2))
        if false {
            _ = len(v4)
        } else {
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint32 = (^(v22))
                v2 -= -1
                v6 -= 42
                _ = fi23
                _ = v24
            }
        }
        _ = fi21
        _ = v22
    }
    var v25 bool = false
    switch "\n\t" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        if true {
            var v26 uint16 = 2
            _ = v26
        }
    }
    fmt.Println(v1, v5, v1)
    var v27 uint32 = 2147483647
    var v28 float32 = (float32(1e10) + 0.1)
    var v29 uint64 = v1
    v25 = ((0.0 - -2.25) > 1.5)
    if true {
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 bool = ((v2 ^ (v2 >> ((v3) & 63))) == (v2 >> ((42) & 63)))
            for fi32 := 0; fi32 < 3; fi32++ {
                switch (fi30 % ((-9223372036854775808) | 1)) {
                case 9223372036854775807:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case -9223372036854775808:
                    v28 = ((((v28 + ((1e10 + float32((1.5 - 0.0))) + (v28 + float32(1e10)))) * 1.5) / v28) - v28)
                case 32767:
                    var v33 float32 = v28
                    _ = v33
                }
                fmt.Println(v5, v6)
                _ = fi32
            }
            v34 := GenS2{F0: -1, F1: 42}
            _ = fi30
            _ = v31
            _ = v34
        }
    }
    v5 += 1
    v35 := GenS3{F0: v27, F1: v2, F2: 1}
    if v25 {
        var v36 uint8 = (v5 - (v5 ^ (^(v5))))
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
