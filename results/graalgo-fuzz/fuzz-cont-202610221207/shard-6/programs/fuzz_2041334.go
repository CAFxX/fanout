package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint, p1 uint8, p2 int32) int32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint8 = p1
        if false {
            var v6 uint32 = 2147483648
            _ = v6
        } else {
            if false {
                var v7 uint = (p0 >> ((p0) & 63))
                v8 := []int{(-(fi4)), (fi4 ^ fi4), (fi4 % (((fi4 + (fi4 + 32767))) | 1)), (fi4 >> ((p0) & 63))}
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := GenS1{F0: (("" + "") + "\n\t"), F1: 0, F2: 2}
                    v11 := []int{fi9, fi9, (fi9 / (((fi4 - fi9)) | 1)), (fi9 ^ 9223372036854775807)}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 uint8 = p1
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(fi13)
            var v14 int64 = -1
            _ = fi13
            _ = v14
        }
        _ = fi4
        _ = v5
    }
    var v15 float32 = (1.5 / -2.25)
    v16 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return p2
}

func main() {
    var v1 int32 = 42
    var v2 float32 = ((0.1 + 1.5) - float32(-2.25))
    var v3 int64 = 719
    var v17 uint64 = 42
    switch "hello" {
    case "":
        v18 := GenS1{F0: "a", F1: v3, F2: (v17 >> ((0) & 63))}
        _ = v18
    case "foo bar":
        var v19 uint16 = 537
        _ = v19
    }
    switch ("\n\t" + "") {
    case "\n\t":
        if ((false && true) && false) {
            v1 -= v1
            fmt.Println(v2, v1, v1)
        }
    case "foo bar":
        v20 := GenS1{F0: "", F1: -1, F2: (^(v17))}
        _ = v20
    case "":
        switch int(v1) {
        case 1:
            v1 -= (v1 ^ (v1 >> ((605) & 63)))
        default:
            v21 := []int{0, 1}
            _ = v21
        }
    }
    switch 395 {
    case -526:
        v22 := GenS3{F0: 42}
        _ = v22
    case -213:
        v1 += (v1 | v1)
    default:
        switch "foo bar" {
        case "foo bar":
            if false {
                v23 := GenS3{F0: (-(v3))}
                v17 = v17
                _ = v23
            }
        case "":
            var v24 float32 = v2
            _ = v24
        default:
            v2 += (v2 + (((-2.25 - -2.25) - v2) - v2))
        }
    }
    fmt.Println(v3, v17)
    v1 = genFn0(2, 74, 0)
    switch -557 {
    case -1:
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 float64 = float64(-2.25)
                if ("a" < ("" + "foo bar")) {
                    fmt.Println(v2)
                    var v28 float64 = (v27 - (-2.25 + (v27 * (float64(v2) - float64(((v2 + v2) * float32(v27)))))))
                    _ = v28
                } else {
                    switch 4294967295 {
                    case -2147483648:
                        fmt.Println(v1)
                    case 32767:
                        var v29 int = fi25
                        _ = v29
                    case 42:
                        var v30 string = "\n\t"
                        _ = v30
                    default:
                        var v31 int32 = -1
                        _ = v31
                    }
                    v27 += float64(v2)
                    v32 := GenS3{F0: (v3 >> ((4294967295) & 63))}
                    _ = v32
                }
                v33 := GenS4{F0: 65535, F1: -477}
                fmt.Println(v2, fi26, v1)
                _ = fi26
                _ = v27
                _ = v33
            }
            fmt.Println(v1, v3)
            v34 := []int{}
            _ = fi25
            _ = v34
        }
    case -2147483648:
        switch "foo bar" {
        case "hello":
            v35 := []int{-71, -282}
            _ = v35
        case "foo bar":
            for fi36 := 0; fi36 < 18; fi36++ {
                v17 -= (v17 + v17)
                _ = fi36
            }
        case "":
            fmt.Println(v1, v1, v17)
        }
    default:
        var v37 uint32 = 65535
        _ = v37
    }
    var v38 int64 = v3
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
}
