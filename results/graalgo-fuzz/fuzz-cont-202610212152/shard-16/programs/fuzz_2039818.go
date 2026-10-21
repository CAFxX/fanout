package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 string }

func main() {
    var v1 int32 = 978
    var v2 bool = false
    var v3 int32 = (v1 + (v1 ^ 1))
    var v4 uint16 = 32767
    var v5 float32 = 0.1
    var v6 bool = v2
    if v6 {
        for fi7 := 0; fi7 < 3; fi7++ {
            switch (fi7 + (^(fi7))) {
            case 21:
                v8 := GenS1{F0: v4}
                _ = v8
            default:
                v9 := GenS1{F0: (v4 / (((v4 & 42)) | 1))}
                _ = v9
            }
            _ = fi7
        }
        if ((^(v4)) < v4) {
            v10 := GenS2{F0: -9223372036854775808}
            _ = v10
        }
        v11 := []int{4294967295, 127, 620, 16}
        _ = v11
    }
    if v2 {
        if (255 < 1) {
            v12 := GenS2{F0: 9223372036854775807}
            _ = v12
        } else {
            v5 += (1.5 + v5)
            fmt.Println(v6, v3, v3)
        }
        if (!true) {
            for fi13 := 0; fi13 < 11; fi13++ {
                fmt.Println(v4, fi13)
                v1 = (v1 | v3)
                _ = fi13
            }
            if true {
                fmt.Println(v3, v4, v4)
                v2 = ((true || (!true)) && true)
                var v14 bool = true
                _ = v14
            }
        }
        var v15 bool = v2
        _ = v15
    } else {
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v6)
            var v17 float32 = (3.14159 + float32((-2.25 / 1.5)))
            _ = fi16
            _ = v17
        }
        v6 = (18446744073709551615 < 1)
    }
    v5 -= -2.25
    var v18 string = ("foo bar" + "\n\t")
    if (18446744073709551615 < 4294967295) {
        v19 := GenS3{F0: 1}
        v20 := []int{149}
        _ = v19
        _ = v20
    } else {
        switch 0 {
        case -1:
            fmt.Println(v3, v5, v4)
        case 9223372036854775807:
            fmt.Println(v4, v18)
        default:
            var v21 uint16 = (v4 + 0)
            _ = v21
        }
        switch 127 {
        case -9223372036854775808:
            v22 := []int{-9223372036854775808}
            _ = v22
        case 2147483647:
            v4 += 0
        }
    }
    var v23 float32 = -2.25
    switch ("\n\t" + "a") {
    case "a":
        var v24 uint32 = 0
        _ = v24
    }
    v25 := []int{375, 2147483647, 128, 384}
    for fi26 := 0; fi26 < 10; fi26++ {
        if false {
            switch v18 {
            case "a":
                if v6 {
                    switch (v18 + "\n\t") {
                    case "":
                        v27 := GenS4{F0: (-(v3))}
                        _ = v27
                    case "\n\t":
                        v18 = v18
                    case "a":
                        var v28 float32 = 0.0
                        _ = v28
                    }
                    v29 := GenS5{F0: 854, F1: (v3 | v3)}
                    _ = v29
                } else {
                    v2 = v6
                    switch v18 {
                    case "foo bar":
                        v25 = append(v25, fi26)
                    }
                    v30 := GenS3{F0: -9223372036854775808}
                    _ = v30
                }
            }
        } else {
            var v31 string = "\n\t"
            _ = v31
        }
        var v32 uint = 381
        _ = fi26
        _ = v32
    }
    var v33 uint64 = 0
    v6 = (!true)
    v34 := GenS6{F0: v18}
    fmt.Println(v6, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
