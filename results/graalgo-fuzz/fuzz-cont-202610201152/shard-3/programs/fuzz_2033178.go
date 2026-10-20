package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int8; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 int8; F2 uint32 }
type GenS5 struct { F0 string; F1 int16; F2 bool }
func genFn0() uint {
    if false {
        if (true || true) {
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: 128, F1: 0}
                _ = fi7
                _ = v8
            }
        }
        var v9 int16 = 762
        if true {
            v9 = v9
        } else {
            var v10 uint16 = 65535
            switch 2147483647 {
            case 42:
                if (3.14159 > 1.5) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v12 := GenS1{F0: 231, F1: 1}
                        v13 := []int{(-(fi11)), (fi11 << ((221) & 63)), fi11}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS2{F0: 42, F1: 4294967295}
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 string = ("\n\t" + "a")
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
            case 4294967295:
                v10 = 65535
            default:
                var v17 string = "\n\t"
                _ = v17
            }
            _ = v10
        }
        _ = v9
    }
    var v18 int8 = 42
    v19 := []int{-267, -457, 9223372036854775807}
    _ = v18
    _ = v19
    return 121
}
func genFn1() bool {
    var v20 int8 = 127
    v20 -= (^(v20))
    if false {
        v21 := []int{3, -375, 127, 127}
        var v22 float64 = 0.0
        var v23 int32 = 0
        _ = v21
        _ = v22
        _ = v23
    } else {
        if true {
            if (false || false) {
                var v24 int64 = -456
                for fi25 := 0; fi25 < 15; fi25++ {
                    v20 -= (v20 % ((-1) | 1))
                    switch fi25 {
                    case -689:
                        v26 := GenS3{F0: (v20 | (-(v20))), F1: v24, F2: 590}
                        _ = v26
                    }
                    _ = fi25
                }
                _ = v24
            } else {
                var v27 uint64 = 2
                _ = v27
            }
        }
        if (-9223372036854775808 < 42) {
            var v28 int = -1
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 uint64 = 465
                if false {
                    v31 := []int{-2147483648, -9223372036854775808, v28}
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            v32 := GenS4{F0: "hello", F1: (v20 | (v20 / (((v20 + (v20 >> ((65535) & 63)))) | 1))), F2: 1}
            _ = v28
            _ = v32
        }
    }
    if true {
        var v33 uint = 327
        v20 = (^(v20))
        switch ("a" + "hello") {
        case "foo bar":
            var v34 uint8 = 128
            _ = v34
        case "hello":
            var v35 int = 32768
            _ = v35
        }
        _ = v33
    } else {
        v36 := GenS1{F0: 2, F1: (^(v20))}
        switch 32768 {
        case 9223372036854775807:
            if (9223372036854775807 < -750) {
                v37 := GenS5{F0: ("" + "foo bar"), F1: 32767, F2: false}
                if true {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 uint64 = 0
                        _ = fi38
                        _ = v39
                    }
                }
                v40 := []int{127, 9223372036854775807, -1}
                _ = v37
                _ = v40
            } else {
                fmt.Println(v20)
            }
        case -492:
            v20 -= 100
        }
        _ = v36
    }
    _ = v20
    return true
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := []int{699}
    var v4 int8 = 42
    var v5 string = "hello"
    var v6 bool = false
    _ = len(v1)
    var v41 float64 = (-2.25 - 3.14159)
    _ = len(v2)
    v4 = v4
    v42 := GenS2{F0: 156, F1: 2}
    switch 128 {
    case 135:
        var v43 bool = v6
        _ = v43
    default:
        v41 += v41
    }
    var v44 int = 42
    v5 = (("" + (v5 + ((v5 + v5) + v5))) + v5)
    if v6 {
        v45 := GenS3{F0: 42, F1: 42, F2: 359}
        if true {
            v44 = (v44 >> ((genFn0()) & 63))
            for fi46 := 0; fi46 < 20; fi46++ {
                v47 := GenS1{F0: 199, F1: 1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v5 = v5
                _ = fi46
                _ = v47
            }
            var v48 uint16 = 42
            _ = v48
        }
        _ = v45
    }
    var v49 uint32 = 158
    v44 = 964
    if (2 != 42) {
        if false {
            var v50 int16 = -32768
            _ = v50
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v49
    _ = v5
    _ = v6
}
