package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    var v1 string = ""
    var v2 bool = true
    var v3 int16 = -22
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := []int{fi4}
        switch v1 {
        case "foo bar":
            v1 = v1
        case "":
            for fi6 := 0; fi6 < 10; fi6++ {
                v3 = 189
                if false {
                    fmt.Println(fi4, fi6, fi4)
                    fmt.Println(fi4, v1, v3)
                    var v7 int16 = v3
                    _ = v7
                } else {
                    v3 = (v3 >> ((255) & 63))
                    v8 := GenS1{F0: 612, F1: int64(v3), F2: 1}
                    var v9 uint16 = 758
                    var v10 string = (v1 + "a")
                    _ = v10
                    _ = v8
                    _ = v9
                }
                _ = fi6
            }
        default:
            for fi11 := 0; fi11 < 14; fi11++ {
                if false {
                    v2 = false
                    v3 -= v3
                    v1 = (v1 + v1)
                    switch v1 {
                    case "a":
                        var v12 int8 = 127
                        _ = v12
                    case "foo bar":
                        var v13 int16 = v3
                        _ = v13
                    default:
                        v5 = append(v5, (fi11 & 9223372036854775807))
                    }
                } else {
                    var v14 int8 = 127
                    if true {
                        v2 = (-1 > -1)
                    }
                    fmt.Println(v2, v3)
                    _ = v14
                }
                var v15 uint16 = 936
                _ = fi11
                _ = v15
            }
        }
        if v2 {
            fmt.Println(v3)
        }
        _ = fi4
        _ = v5
    }
    v16 := GenS1{F0: 2147483648, F1: 42, F2: 210}
    v17 := GenS1{F0: 65535, F1: 1, F2: 42}
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS2{F0: 1}
            var v21 int = (fi19 << ((255) & 63))
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int64 = 1
        v23 := GenS2{F0: 42}
        _ = fi18
        _ = v22
        _ = v23
    }
    var v24 uint32 = 0
    v2 = v2
    fmt.Println(v24, v3, v2)
    fmt.Println(v24, v24)
    if v2 {
        fmt.Println(v3, v24, v24)
        for fi25 := 0; fi25 < 7; fi25++ {
            switch (fi25 - fi25) {
            case 346:
                v26 := GenS1{F0: 634, F1: 1, F2: 0}
                _ = v26
            }
            var v27 uint32 = (v24 ^ 42)
            _ = fi25
            _ = v27
        }
        v1 = v1
    }
    v2 = (!false)
    v2 = ((1.5 / ((((1.5 + -2.25) + 0.0) + 0.0) / 0.1)) != float64(0.1))
    var v28 int = int(v3)
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := []int{}
        if (4294967295 < 244) {
            v3 = (v3 / ((0) | 1))
        } else {
            var v31 int = -2147483648
            for fi32 := 0; fi32 < 5; fi32++ {
                fmt.Println(v28, fi32)
                if len(v30) > 0 {
                    v30[0] = v28
                }
                var v33 int8 = -1
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        if (-217 >= -9223372036854775808) {
            for fi34 := 0; fi34 < 6; fi34++ {
                v3 = v3
                _ = len(v30)
                _ = fi34
            }
        }
        _ = fi29
        _ = v30
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        v24 = 0
        _ = fi35
    }
    v2 = ((!((190 != 2) && (!((v28 % ((v28) | 1)) != v28)))) && v2)
    v36 := GenS3{F0: 890, F1: 255, F2: 0}
    switch v28 {
    case 583:
        v1 = (("a" + ("foo bar" + v1)) + (((("" + v1) + "foo bar") + (("foo bar" + v1) + v1)) + ((v1 + v1) + v1)))
    default:
        v1 = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v36
}
