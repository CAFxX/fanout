package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int }
func genFn0(p0 uint, p1 bool, p2 int64) int8 {
    if p1 {
        v4 := GenS1{F0: (p2 / (((p2 & p2)) | 1))}
        if p1 {
            v5 := []int{127, 32767}
            p1 = false
            _ = v5
        }
        _ = v4
    } else {
        if false {
            v6 := []int{32767, 128}
            if (1.5 == 1.5) {
                v7 := GenS2{F0: 2147483647}
                var v8 int32 = -2147483648
                p1 = true
                _ = v7
                _ = v8
            } else {
                var v9 float32 = 0.0
                _ = v9
            }
            var v10 int32 = 180
            _ = v10
            _ = v6
        } else {
            if (!(false || ((!p1) && ("a" <= "\n\t")))) {
                var v11 bool = true
                fmt.Println(v11)
                _ = v11
            } else {
                p2 -= p2
            }
        }
    }
    p0 += (p0 | (p0 / ((270) | 1)))
    fmt.Println(p0, p0)
    if (-2.25 >= (1.5 - (3.14159 + ((1.5 - float32(-2.25)) + ((1e10 * 0.1) / (3.14159 + 1e10)))))) {
        if (-32768 < 0) {
            switch 127 {
            case 42:
                for fi12 := 0; fi12 < 6; fi12++ {
                    switch "hello" {
                    case "\n\t":
                        p0 = (p0 << ((p0) & 63))
                    case "foo bar":
                        switch (("a" + "") + ("a" + "foo bar")) {
                        case "\n\t":
                            var v13 uint8 = 99
                            _ = v13
                        default:
                            for fi14 := 0; fi14 < 20; fi14++ {
                                var v15 uint16 = 2
                                fmt.Println(p0, p2)
                                v16 := []int{(fi12 ^ (fi12 / ((fi14) | 1))), (fi14 | (fi12 ^ fi14))}
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        }
                    case "a":
                        v17 := GenS1{F0: p2}
                        _ = v17
                    default:
                        fmt.Println(p0)
                    }
                    var v18 uint8 = 255
                    v19 := GenS2{F0: 2147483647}
                    _ = fi12
                    _ = v18
                    _ = v19
                }
            case 9223372036854775807:
                v20 := []int{42, 288}
                _ = v20
            default:
                for fi21 := 0; fi21 < 20; fi21++ {
                    if p1 {
                        var v22 string = ""
                        p2 += p2
                        v23 := []int{fi21}
                        _ = v22
                        _ = v23
                    } else {
                        v24 := []int{fi21, fi21, fi21, 0}
                        _ = v24
                    }
                    v25 := GenS2{F0: 2147483647}
                    p0 -= (p0 | (^(p0)))
                    _ = fi21
                    _ = v25
                }
            }
            var v26 int16 = -322
            _ = v26
        }
    } else {
        fmt.Println(p2)
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                p1 = (255 > p0)
                _ = fi28
            }
            v29 := []int{fi27, (fi27 - (fi27 * (^(fi27))))}
            _ = fi27
            _ = v29
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1(p0 uint) int64 {
    v30 := GenS1{F0: -9223372036854775808}
    var v31 string = "\n\t"
    v32 := GenS3{F0: -2147483648, F1: 42}
    var v33 int8 = 127
    _ = p0
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    return -203
}

func main() {
    v1 := []string{"a"}
    var v2 uint8 = 0
    var v3 uint8 = 2
    v34 := GenS4{F0: 0, F1: 2}
    if false {
        v2 = (v3 + 4)
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 6; fi36++ {
                v2 -= (v3 - v2)
                v2 = 24
                v37 := GenS5{F0: -332}
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
    }
    v2 -= 42
    v1 = append(v1, ("a" + "a"))
    _ = len(v1)
    for fi38 := 0; fi38 < 5; fi38++ {
        v39 := GenS4{F0: 200, F1: 2147483648}
        var v40 float32 = 3.14159
        v3 += 128
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS5{F0: 0}
    v3 += v2
    fmt.Println(v2, v2)
    if true {
        var v42 int64 = -1
        v3 += v3
        _ = v42
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        _ = len(v1)
        fmt.Println(v3)
        _ = fi43
    }
    var v44 int16 = 1
    var v45 int16 = (v44 ^ (v44 + -32768))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v41
    _ = v44
    _ = v45
}
