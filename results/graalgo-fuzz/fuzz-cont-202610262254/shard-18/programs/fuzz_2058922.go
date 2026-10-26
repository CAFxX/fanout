package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32 }
func genFn0() uint {
    var v4 uint32 = 1
    v5 := GenS1{F0: (v4 | 529), F1: -1}
    v4 -= (^(v4))
    v4 = 65535
    v6 := []int{-9223372036854775808, 1}
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 bool = true
    var v2 string = ("hello" + "a")
    v3 := make(map[uint32]bool)
    v1 = (1 <= -128)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int = 0
        _ = fi7
        _ = v8
    }
    if (387 == 2147483648) {
        v9 := []int{}
        var v10 uint = 18446744073709551615
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS2{F0: "hello"}
                v1 = v1
                fmt.Println(v1)
                if (!(464 >= int64(fi11))) {
                    if false {
                        v10 = 945
                    }
                    switch fi12 {
                    case 127:
                        fmt.Println(fi11)
                    case 32767:
                        if false {
                            v14 := GenS3{F0: true}
                            _ = v14
                        } else {
                            switch (v2 + v2) {
                            case "a":
                                v2 = v2
                            }
                            v1 = true
                        }
                    case 4294967295:
                        var v15 string = (((v2 + "hello") + v2) + v2)
                        _ = v15
                    default:
                        var v16 uint16 = 528
                        _ = v16
                    }
                    v1 = (42 >= 1)
                } else {
                    var v17 int8 = 1
                    v17 = (v17 << (((v10 >> ((255) & 63))) & 63))
                    var v18 bool = v1
                    _ = v17
                    _ = v18
                }
                v10 -= (v10 >> ((v10) & 63))
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    v19 := v3[321]
    v19 = v1
    var v20 int16 = 0
    switch 1 {
    case 175:
        v21 := []int{}
        _ = v21
    case 128:
        switch (v2 + v2) {
        case "foo bar":
            v20 -= v20
        default:
            if v1 {
                v19 = v19
                var v22 int64 = 42
                _ = v22
            }
        }
    case 42:
        var v23 float32 = 3.14159
        _ = v23
    }
    v24 := GenS3{F0: false}
    v19 = v19
    var v25 uint8 = 255
    if false {
        var v26 int8 = 42
        v25 = v25
        for fi27 := 0; fi27 < 18; fi27++ {
            v25 = v25
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 float64 = (-2.25 / 0.1)
                _ = len(v3)
                if (42 < -2147483648) {
                    if (42 > 698) {
                        v30 := GenS3{F0: v1}
                        v29 -= float64(0.1)
                        var v31 int = fi27
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v32 bool = false
                    switch 2147483647 {
                    case -265:
                        if (0 > -9223372036854775808) {
                            v20 = v20
                            if v32 {
                                v32 = v19
                                v32 = (-1 != 584)
                                var v33 string = v2
                                _ = v33
                            }
                            v34 := []int{fi27}
                            v32 = (603 > 255)
                            if true {
                                v1 = v19
                            }
                            _ = v34
                        } else {
                            v35 := GenS4{F0: 0}
                            _ = v35
                        }
                    case -1:
                        v20 += (v20 >> ((18446744073709551615) & 63))
                        switch fi27 {
                        case 32768:
                            var v36 uint64 = 4294967295
                            _ = v36
                        case 128:
                            v37 := GenS5{F0: v2}
                            _ = v37
                        case -2147483648:
                            v38 := []int{(fi28 - (^(fi27)))}
                            _ = v38
                        default:
                            v39 := GenS1{F0: 42, F1: (v20 * v20)}
                            _ = v39
                        }
                    default:
                        v40 := GenS5{F0: ("" + ((v2 + "\n\t") + v2))}
                        _ = v40
                    }
                    v20 = (v20 ^ (v20 & (v20 / ((v20) | 1))))
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
            if v19 {
                var v41 uint = 0
                _ = v41
            }
            _ = fi27
        }
        _ = v26
    }
    _ = len(v3)
    v42 := GenS6{F0: 883}
    v43 := GenS1{F0: 619, F1: 889}
    for fi44 := 0; fi44 < 17; fi44++ {
        delete(v3, 723)
        _ = fi44
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        if v19 {
            v45 := GenS5{F0: ("hello" + v2)}
            _ = v45
        } else {
            delete(v3, 1)
        }
    case "":
        switch (("\n\t" + "a") + v2) {
        case "foo bar":
            switch 522 {
            case 42:
                for fi46 := 0; fi46 < 6; fi46++ {
                    var v47 uint16 = 2
                    _ = fi46
                    _ = v47
                }
            case 525:
                v2 = "a"
            }
        case "hello":
            for fi48 := 0; fi48 < 5; fi48++ {
                switch (fi48 + (fi48 % ((fi48) | 1))) {
                case 128:
                    switch fi48 {
                    case 42:
                        var v49 uint32 = 445
                        _ = v49
                    case 9223372036854775807:
                        v2 = ((v2 + ("a" + v2)) + v2)
                    }
                case 2147483647:
                    fmt.Println(v19, fi48, v25)
                case 32768:
                    var v50 string = (v2 + v2)
                    _ = v50
                }
                _ = fi48
            }
        default:
            _ = len(v3)
        }
    case "a":
        v51 := GenS4{F0: 438}
        _ = v51
    default:
        if (!(!(v20 < (v20 >> ((4294967295) & 63))))) {
            var v52 uint8 = 2
            v19 = (85 != -54)
            switch 0 {
            case 321:
                switch "foo bar" {
                case "\n\t":
                    v1 = (!true)
                }
            case 1:
                var v53 int16 = (v20 >> ((65535) & 63))
                _ = v53
            default:
                fmt.Println(v2, v20, v20)
            }
            _ = v52
        }
    }
    v54 := []int{-1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v42
    _ = v43
    _ = v54
}
