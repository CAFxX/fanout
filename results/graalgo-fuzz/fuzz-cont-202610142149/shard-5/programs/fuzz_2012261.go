package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = -32768
    var v3 bool = ((466 < 0) || true)
    v4 := make(map[string]uint32)
    v2 += (v2 - -566)
    for fi5 := 0; fi5 < 18; fi5++ {
        if ((float64((1.5 / -2.25)) <= 0.0) || v3) {
            var v6 int16 = 535
            _ = v6
        }
        v7 := GenS1{F0: 42}
        var v8 uint64 = 68
        _ = fi5
        _ = v7
        _ = v8
    }
    switch -824 {
    case 447:
        v2 = v2
    case 128:
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 51}
            v1 = v1
            if (2 >= 1) {
                v4[(("hello" + "foo bar") + "")] = 673
                v11 := GenS1{F0: 0}
                _ = v11
            } else {
                var v12 int16 = (v2 << ((2) & 63))
                if false {
                    v13 := GenS1{F0: 1}
                    v14 := GenS1{F0: 42}
                    _ = v13
                    _ = v14
                } else {
                    if v3 {
                        fmt.Println(fi9, v3)
                        v2 = v2
                    }
                }
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    default:
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 float32 = 1.5
                v3 = v3
                switch (^(fi15)) {
                case 176:
                    v3 = true
                case -333:
                    v2 = 502
                }
                _ = fi16
                _ = v17
            }
            switch "a" {
            case "\n\t":
                v18 := GenS1{F0: 0}
                _ = v18
            default:
                fmt.Println(v3, fi15, v3)
            }
            v19 := []int{(^(fi15)), (fi15 * (fi15 >> ((135) & 63))), (fi15 ^ fi15)}
            _ = fi15
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 uint32 = 2147483647
            var v23 uint16 = 255
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 int = fi20
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS1{F0: 1}
            _ = fi25
            _ = v26
        }
        _ = fi20
        _ = v24
    }
    if v3 {
        var v27 float32 = 0.0
        v28 := v4["\n\t"]
        fmt.Println(v28, v27)
        _ = v27
        _ = v28
    }
    var v29 int32 = v1
    for fi30 := 0; fi30 < 19; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            if (!v3) {
                v32 := v4[""]
                var v33 uint32 = 1
                _ = v32
                _ = v33
            } else {
                if ((0.1 + float32((3.14159 + (float64(1e10) + 0.1)))) != 0.1) {
                    v3 = (!(true || (-2.25 <= 0.1)))
                } else {
                    var v34 int32 = -653
                    var v35 uint16 = 981
                    _ = v34
                    _ = v35
                }
                v2 = v2
            }
            _ = fi31
        }
        _ = fi30
    }
    fmt.Println(v3, v1)
    v29 += v29
    if (v3 || (v3 && v3)) {
        if v3 {
            var v36 uint64 = 42
            _ = v36
        }
        switch -1 {
        case 9223372036854775807:
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 int16 = (v2 ^ (^(v2)))
                delete(v4, "\n\t")
                _ = fi37
                _ = v38
            }
        case 42:
            switch ("" + "hello") {
            case "":
                var v39 uint = 42
                _ = v39
            case "a":
                _ = len(v4)
            }
        default:
            var v40 int = 2147483647
            _ = v40
        }
        var v41 uint32 = 248
        _ = v41
    }
    v42 := []int{-519, -9223372036854775808, -873}
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 uint32 = 100
        if true {
            var v45 uint16 = 26
            _ = v45
        }
        var v46 uint64 = 0
        _ = fi43
        _ = v44
        _ = v46
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        v48 := GenS1{F0: 76}
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        var v50 float32 = 0.1
        var v51 float64 = (float64((float32(float64((v50 - 0.1))) * v50)) - 1.5)
        v2 += (v2 ^ (v2 << ((42) & 63)))
        _ = fi49
        _ = v50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v42
}
