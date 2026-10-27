package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
func genFn0() int16 {
    var v5 int64 = 1
    v6 := GenS1{F0: true}
    _ = v5
    _ = v6
    return 1000
}
func genFn1() uint16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 float32 = (0.1 - float32((-2.25 / -2.25)))
            if false {
                var v10 float32 = v9
                v11 := []int{fi8, (fi8 ^ fi7), fi8}
                v11 = append(v11, fi8)
                v12 := GenS1{F0: (32767 != 32767)}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        var v13 int8 = 4
        _ = fi7
        _ = v13
    }
    v14 := GenS1{F0: (29 == 1)}
    for fi15 := 0; fi15 < 16; fi15++ {
        if false {
            var v16 int = -323
            switch (v16 | fi15) {
            case -443:
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := []int{v16, (-(fi15)), v16}
                    v19 := GenS1{F0: (!(25 <= -79))}
                    var v20 int = (fi15 | fi17)
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            default:
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 int = 9223372036854775807
                    switch -937 {
                    case 493:
                        var v23 float64 = (1e10 * 1.5)
                        _ = v23
                    }
                    v22 += 0
                    _ = fi21
                    _ = v22
                }
            }
            v24 := []int{v16, v16}
            _ = v16
            _ = v24
        }
        v25 := GenS1{F0: (genFn0() != -465)}
        _ = fi15
        _ = v25
    }
    _ = v14
    return 577
}

func main() {
    v1 := []int{}
    var v2 uint16 = 776
    var v3 bool = false
    var v4 int16 = 344
    v4 = genFn0()
    if v3 {
        if true {
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v3, fi26)
                v27 := GenS1{F0: true}
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                v2 += (v2 ^ (v2 ^ (v2 >> ((947) & 63))))
                _ = fi28
            }
            if len(v1) > 0 {
                v1[0] = 977
            }
        } else {
            var v29 int64 = -1
            var v30 uint32 = 4294967295
            _ = v29
            _ = v30
        }
        var v31 float32 = 0.0
        _ = v31
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v32 := GenS1{F0: true}
        var v33 uint8 = 205
        _ = v32
        _ = v33
    } else {
        fmt.Println(v2, v2)
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        _ = len(v1)
        v3 = v3
        switch (fi34 % (((fi34 % (((fi34 % ((2147483647) | 1))) | 1))) | 1)) {
        case 128:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 0:
            v2 += v2
        default:
            v4 -= 1
        }
        _ = fi34
    }
    var v35 uint64 = 42
    v36 := GenS1{F0: (((1.5 == 3.14159) || v3) && (-2.25 == float32(-2.25)))}
    for fi37 := 0; fi37 < 17; fi37++ {
        v2 += (v2 * 405)
        var v38 uint64 = v35
        _ = fi37
        _ = v38
    }
    v39 := GenS1{F0: (!("\n\t" > "hello"))}
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v3)
    case "a":
        v40 := GenS2{F0: int32(v35)}
        _ = v40
    default:
        var v41 bool = true
        _ = v41
    }
    v42 := GenS2{F0: -1}
    switch (((("hello" + "\n\t") + "a") + "") + "foo bar") {
    case "\n\t":
        switch ("foo bar" + (("hello" + "\n\t") + ("hello" + ""))) {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        switch ("a" + "foo bar") {
        case "foo bar":
            if true {
                v2 = v2
                v4 = (v4 / (((v4 & genFn0())) | 1))
            } else {
                v2 = (v2 % ((v2) | 1))
                v2 += 32767
            }
        default:
            fmt.Println(v4)
        }
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 int8 = 100
        _ = fi43
        _ = v44
    }
    v45 := GenS2{F0: 42}
    for fi46 := 0; fi46 < 1; fi46++ {
        fmt.Println(v3, v2)
        _ = fi46
    }
    v3 = v3
    var v47 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v42
    _ = v45
    _ = v47
}
