package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int8 }
func genFn0() int {
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    v8 := GenS1{F0: 4294967295}
    if (!true) {
        if ((true && (2 <= 42)) || (0 != 42)) {
            switch 0 {
            case 127:
                if true {
                    var v9 uint64 = 559
                    if ((!(!false)) || (53 < 896)) {
                        if (!true) {
                            var v10 uint = 457
                            _ = v10
                        }
                        v11 := GenS2{F0: 688, F1: 2}
                        _ = v11
                    } else {
                        switch 4294967295 {
                        case 520:
                            var v12 string = "foo bar"
                            _ = v12
                        case 0:
                            v13 := []int{32767}
                            _ = v13
                        case 754:
                            switch 42 {
                            case 32767:
                                if false {
                                    var v14 int64 = 1
                                    _ = v14
                                } else {
                                    v9 += 2
                                }
                            case 9223372036854775807:
                                if true {
                                    var v15 string = "a"
                                    _ = v15
                                }
                            case 42:
                                if (!(2 <= 1)) {
                                    v16 := GenS3{F0: "a", F1: 68}
                                    v9 += (v9 >> ((18446744073709551615) & 63))
                                    _ = v16
                                }
                            }
                        default:
                            v9 = v9
                        }
                    }
                    _ = v9
                } else {
                    var v17 int = -717
                    v18 := []int{}
                    _ = v17
                    _ = v18
                }
            case -1:
                var v19 bool = false
                _ = v19
            case -9223372036854775808:
                for fi20 := 0; fi20 < 5; fi20++ {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 int8 = 1
                        _ = fi21
                        _ = v22
                    }
                    _ = fi20
                }
            }
        } else {
            var v23 int32 = -2147483648
            v24 := GenS2{F0: 9223372036854775807, F1: 0}
            var v25 uint = 255
            _ = v23
            _ = v24
            _ = v25
        }
    } else {
        var v26 string = ""
        v27 := []int{128, 127}
        _ = v26
        _ = v27
    }
    _ = v8
    return -1
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]string)
    var v3 uint = 1
    v4 := make(map[string]int)
    var v5 int64 = 1
    v6 := make(map[string]uint32)
    if false {
        for fi28 := 0; fi28 < 7; fi28++ {
            if true {
                v6[("a" + "foo bar")] = 0
                v4[("" + "\n\t")] = 127
            }
            switch genFn0() {
            case 42:
                for fi29 := 0; fi29 < 10; fi29++ {
                    v3 -= v3
                    v30 := GenS2{F0: v5, F1: 128}
                    if (486 > (v3 % ((v3) | 1))) {
                        var v31 uint = v3
                        _ = v31
                    }
                    _ = fi29
                    _ = v30
                }
            case 128:
                delete(v4, "foo bar")
            }
            _ = fi28
        }
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint32 = 393
            _ = fi32
            _ = v33
        }
        fmt.Println(v3, v5)
    }
    if true {
        switch (v1 ^ v1) {
        case 0:
            v3 = (v3 + (v3 % ((v3) | 1)))
        }
    }
    if (!false) {
        v34 := []int{927, (-(v1)), (v1 * v1)}
        delete(v2, (v1 ^ (v1 << ((659) & 63))))
        _ = v34
    } else {
        v35 := GenS2{F0: 0, F1: 128}
        if true {
            var v36 int64 = -200
            _ = v36
        } else {
            v37 := GenS1{F0: 261}
            switch 1 {
            case 0:
                v38 := []int{(v1 ^ v1), genFn0(), genFn0()}
                _ = v38
            case -2147483648:
                if true {
                    v39 := GenS4{F0: 0}
                    switch (v1 / ((1) | 1)) {
                    case 1:
                        fmt.Println(v3, v5, v5)
                    case 269:
                        v5 -= v5
                    default:
                        for fi40 := 0; fi40 < 10; fi40++ {
                            var v41 int32 = 0
                            v42 := GenS3{F0: "\n\t", F1: 127}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    }
                    for fi43 := 0; fi43 < 15; fi43++ {
                        fmt.Println(v5)
                        _ = fi43
                    }
                    _ = v39
                }
            case -492:
                v3 = 255
            default:
                v1 = v1
            }
            _ = v37
        }
        _ = v35
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        v1 = 0
        var v45 bool = false
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 uint32 = 65535
            var v48 uint16 = 1
            v45 = (!true)
            fmt.Println(v1, v5)
            v3 = v3
            delete(v2, genFn0())
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = fi44
        _ = v45
    }
    for fi49 := 0; fi49 < 18; fi49++ {
        var v50 int64 = v5
        var v51 int = v1
        fmt.Println(v5, v50)
        _ = fi49
        _ = v50
        _ = v51
    }
    v52 := GenS1{F0: 2}
    v3 = (v3 - 4294967295)
    v3 += (v3 + v3)
    for fi53 := 0; fi53 < 11; fi53++ {
        if false {
            for fi54 := 0; fi54 < 18; fi54++ {
                v5 += 381
                var v55 uint16 = 1
                _ = fi54
                _ = v55
            }
            switch 4294967295 {
            case 2147483647:
                var v56 uint16 = 32767
                _ = v56
            case 1:
                for fi57 := 0; fi57 < 17; fi57++ {
                    v58 := GenS2{F0: (v5 | (-(v5))), F1: 1}
                    v1 -= 164
                    _ = fi57
                    _ = v58
                }
            }
        }
        if true {
            v1 = fi53
        }
        _ = fi53
    }
    v5 = 1
    v5 = (v5 ^ (v5 / ((v5) | 1)))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
