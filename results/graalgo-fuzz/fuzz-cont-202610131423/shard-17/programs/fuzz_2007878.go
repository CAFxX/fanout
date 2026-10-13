package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) int32 {
    var v5 string = ("foo bar" + "foo bar")
    var v6 float64 = (3.14159 + (0.0 / 1.0))
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: true}
        v9 := []int{(fi7 - 2147483647)}
        v5 = (v5 + "a")
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{4294967295}
    _ = p0
    _ = v10
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int64 = -474
    var v2 uint16 = 255
    var v3 float64 = (-2.25 * -2.25)
    var v4 bool = ((("hello" + "a") >= "") && (!(29 > 700)))
    var v11 uint16 = v2
    v2 -= (v11 - v11)
    switch ("" + "\n\t") {
    case "foo bar":
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS1{F0: ((fi12 & 595) > 2147483647)}
                var v15 string = ("a" + "a")
                fmt.Println(v11, v15)
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    default:
        v2 = v2
    }
    var v16 int = -2147483648
    var v17 uint = 1
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: ((v2 % (((v2 + (v11 ^ v2))) | 1)) >= v2)}
        v3 = (0.1 - 0.0)
        v3 -= 0.0
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        switch "foo bar" {
        case "":
            if v4 {
                if ((!v4) || v4) {
                    var v21 int32 = genFn0(121)
                    _ = v21
                } else {
                    v22 := GenS1{F0: (-327 > v1)}
                    var v23 int64 = v1
                    _ = v22
                    _ = v23
                }
                var v24 int16 = 345
                if v4 {
                    v11 -= 2
                    v25 := GenS1{F0: v4}
                    var v26 int32 = 0
                    _ = v25
                    _ = v26
                } else {
                    var v27 uint64 = 4294967295
                    for fi28 := 0; fi28 < 6; fi28++ {
                        v27 = (v27 - (^(v27)))
                        fmt.Println(v1, v1)
                        switch (fi28 - (fi28 - (fi20 >> ((255) & 63)))) {
                        case 9223372036854775807:
                            v29 := GenS2{F0: (v16 % ((int(v17)) | 1))}
                            _ = v29
                        }
                        _ = fi28
                    }
                    var v30 uint = v17
                    _ = v27
                    _ = v30
                }
                _ = v24
            } else {
                v31 := GenS3{F0: -183}
                _ = v31
            }
        case "hello":
            switch (fi20 * fi20) {
            case 9223372036854775807:
                if (false && (!(!(4294967295 <= 488)))) {
                    switch "a" {
                    case "":
                        if v4 {
                            v16 = v16
                        } else {
                            var v32 uint8 = 255
                            _ = v32
                        }
                    default:
                        var v33 uint = 0
                        _ = v33
                    }
                    var v34 int64 = v1
                    _ = v34
                }
            }
        }
        _ = fi20
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v16 = (v16 & (-(v16)))
        _ = fi35
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := GenS3{F0: genFn0(4294967295)}
        v1 = (v1 & v1)
        _ = fi36
        _ = v37
    }
    v2 += v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
