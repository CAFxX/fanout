package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0() int64 {
    switch 32767 {
    case -1:
        v4 := []int{2147483647, 0, -1, 360}
        _ = v4
    case -2147483648:
        var v5 uint16 = 0
        _ = v5
    }
    var v6 int32 = 0
    switch "hello" {
    case "\n\t":
        v6 += v6
    }
    v7 := []int{42, -1, -9223372036854775808, 1}
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int8 = 127
    var v2 uint32 = 1
    var v3 uint64 = 1
    var v8 string = "foo bar"
    var v9 string = v8
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{(^(fi10))}
        v9 = (v8 + (v9 + v9))
        v2 = (v2 + (v2 << ((2) & 63)))
        _ = fi10
        _ = v11
    }
    switch 32768 {
    case 42:
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := []int{32768, -2147483648, (fi12 + (fi12 >> ((18446744073709551615) & 63))), fi12}
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint8 = 42
                if true {
                    v1 += v1
                }
                v2 = (^(v2))
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    case 32767:
        for fi16 := 0; fi16 < 9; fi16++ {
            switch 32768 {
            case -81:
                if (true || true) {
                    if (1000 < -32768) {
                        switch 439 {
                        case 1:
                            if (1.5 <= (-2.25 + ((0.1 + float64(3.14159)) / 3.14159))) {
                                v17 := GenS1{F0: v2}
                                var v18 float64 = 0.0
                                switch (fi16 + (fi16 % (((fi16 - (fi16 >> ((2) & 63)))) | 1))) {
                                case -9223372036854775808:
                                    for fi19 := 0; fi19 < 16; fi19++ {
                                        var v20 int8 = v1
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                                _ = v17
                                _ = v18
                            }
                        case 128:
                            for fi21 := 0; fi21 < 19; fi21++ {
                                v9 = "a"
                                v22 := []int{(fi16 & 1)}
                                fmt.Println(v3)
                                var v23 float32 = (float32(1.5) + (3.14159 - 0.0))
                                if len(v22) > 0 {
                                    v22[0] = (fi21 + fi16)
                                }
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        case -618:
                            for fi24 := 0; fi24 < 19; fi24++ {
                                v3 -= (v3 * 18446744073709551615)
                                v25 := []int{}
                                _ = fi24
                                _ = v25
                            }
                        }
                    }
                    if (42 != 42) {
                        v8 = v8
                        v9 = (v8 + (("a" + "") + v9))
                        v26 := []int{(-(fi16)), (fi16 - (fi16 << ((588) & 63))), 602, fi16}
                        _ = v26
                    }
                    var v27 int8 = v1
                    _ = v27
                }
            case 611:
                v28 := []int{fi16, fi16, 0}
                _ = v28
            default:
                if (v1 > (v1 - v1)) {
                    fmt.Println(fi16, fi16, v8)
                }
            }
            _ = fi16
        }
    }
    var v29 uint32 = (v2 / (((v2 + (v2 >> ((324) & 63)))) | 1))
    fmt.Println(v3, v9, v29)
    v8 = (v9 + "")
    v30 := GenS2{F0: -2147483648, F1: 255, F2: 0}
    for fi31 := 0; fi31 < 4; fi31++ {
        if (true || true) {
            v3 = (v3 >> ((42) & 63))
        } else {
            if (!(!false)) {
                var v32 int16 = -32768
                for fi33 := 0; fi33 < 4; fi33++ {
                    v1 += 127
                    _ = fi33
                }
                if false {
                    var v34 int = 128
                    var v35 uint32 = 424
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint8 = 42
                    _ = v36
                }
                _ = v32
            }
        }
        for fi37 := 0; fi37 < 5; fi37++ {
            var v38 uint = 596
            v39 := GenS1{F0: v2}
            v40 := GenS1{F0: v2}
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        _ = fi31
    }
    v41 := GenS1{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v41
    _ = v8
    _ = v9
}
