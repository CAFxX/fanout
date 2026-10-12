package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0() int {
    var v4 bool = (255 == 18446744073709551615)
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 bool = v4
        if v6 {
            v4 = v6
            v6 = false
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint32 = 2147483648
            var v9 int8 = 0
            fmt.Println(v6, fi5, fi7)
            v10 := []int{-709, fi5}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -744
    var v3 uint16 = 357
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint64 = 319
                v3 -= 0
                v14 := v1[604]
                _ = fi12
                _ = v13
                _ = v14
            }
            if false {
                v3 += 255
            } else {
                v3 = v3
                var v15 float32 = 0.0
                _ = v15
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: 547, F1: true, F2: -30}
                if true {
                    fmt.Println(fi11)
                } else {
                    v3 -= 1
                    v18 := []int{}
                    var v19 int32 = (v2 + 745)
                    _ = v18
                    _ = v19
                }
                var v20 float64 = 0.1
                _ = fi16
                _ = v17
                _ = v20
            }
            _ = fi11
        }
        var v21 int32 = v2
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 string = ("" + (("hello" + "\n\t") + "a"))
            if (!false) {
                var v24 int64 = 0
                if true {
                    v23 = v23
                    if (true && false) {
                        if (818 > 255) {
                            for fi25 := 0; fi25 < 2; fi25++ {
                                v1[2147483647] = 42
                                v26 := GenS2{F0: v24, F1: 521}
                                var v27 float32 = ((1e10 - 1e10) * (1.5 - float32(((float64(0.1) - (1.5 * float64(float32(float64(1e10))))) * 1e10))))
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            var v28 uint8 = 254
                            var v29 float32 = 1.5
                            _ = v28
                            _ = v29
                        } else {
                            if true {
                                v24 = 438
                            } else {
                                fmt.Println(v23, v24)
                                switch ((("a" + v23) + (v23 + ((v23 + v23) + (v23 + v23)))) + v23) {
                                case "a":
                                    if false {
                                        var v30 uint8 = 2
                                        _ = v30
                                    } else {
                                        var v31 int32 = 0
                                        switch (fi22 >> ((67) & 63)) {
                                        case 809:
                                            var v32 uint8 = uint8(v3)
                                            _ = v32
                                        default:
                                            fmt.Println(v21, v2, v3)
                                        }
                                        _ = v31
                                    }
                                }
                                delete(v1, 4294967295)
                            }
                        }
                    }
                    var v33 int = fi22
                    _ = v33
                }
                switch 4294967295 {
                case 9223372036854775807:
                    var v34 uint32 = 2
                    _ = v34
                case 32767:
                    switch v23 {
                    case "foo bar":
                        for fi35 := 0; fi35 < 10; fi35++ {
                            if (false || ((^(v3)) == (v3 ^ v3))) {
                                v2 += v2
                                fmt.Println(v3, v3)
                                if (!false) {
                                    var v36 bool = (123 == 2)
                                    _ = v36
                                }
                            }
                            var v37 string = ((v23 + (v23 + "\n\t")) + v23)
                            if false {
                                switch (fi22 - genFn0()) {
                                case 127:
                                    var v38 uint32 = 460
                                    _ = v38
                                case -1:
                                    var v39 uint64 = 662
                                    _ = v39
                                }
                                if false {
                                    v24 -= (v24 + int64(v3))
                                    v37 = v37
                                    fmt.Println(v23)
                                }
                                v3 = v3
                            } else {
                                switch ((("\n\t" + ((v37 + "a") + ("" + ((v37 + "foo bar") + "hello")))) + v23) + v23) {
                                case "hello":
                                    v40 := v1[2147483647]
                                    _ = v40
                                case "":
                                    var v41 uint = 65535
                                    fmt.Println(v41)
                                    _ = v41
                                }
                                v24 = (v24 << ((793) & 63))
                                fmt.Println(v23, v3, v2)
                            }
                            _ = fi35
                            _ = v37
                        }
                    default:
                        var v42 int32 = (v21 & (v2 & int32(fi22)))
                        _ = v42
                    }
                case 42:
                    for fi43 := 0; fi43 < 2; fi43++ {
                        switch v23 {
                        case "":
                            var v44 int8 = -1
                            _ = v44
                        default:
                            v45 := GenS3{F0: 580, F1: 255, F2: v24}
                            _ = v45
                        }
                        delete(v1, 2147483648)
                        _ = fi43
                    }
                }
                _ = v24
            }
            v3 -= (v3 * v3)
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        var v46 uint8 = 250
        v2 -= 140
        _ = v46
    }
    var v47 string = ("" + "\n\t")
    if ((v2 >> ((885) & 63)) == 844) {
        if (true || false) {
            delete(v1, 1)
        }
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 float32 = 0.1
            _ = fi48
            _ = v49
        }
    } else {
        for fi50 := 0; fi50 < 16; fi50++ {
            v2 = v2
            _ = fi50
        }
        for fi51 := 0; fi51 < 8; fi51++ {
            v52 := GenS2{F0: -9223372036854775808, F1: 2}
            _ = fi51
            _ = v52
        }
    }
    if (!true) {
        switch 9223372036854775807 {
        case 2147483647:
            for fi53 := 0; fi53 < 6; fi53++ {
                var v54 uint16 = (v3 | v3)
                v2 -= 0
                _ = fi53
                _ = v54
            }
        default:
            v55 := GenS3{F0: -477, F1: 255, F2: -1}
            _ = v55
        }
        if (true || true) {
            v56 := GenS4{F0: 4294967295, F1: (false || false)}
            v57 := []int{genFn0(), 4294967295}
            _ = v56
            _ = v57
        } else {
            v3 = v3
        }
    } else {
        var v58 int64 = -339
        _ = v58
    }
    fmt.Println(v47, v47, v47)
    v47 = v47
    var v59 float64 = 0.0
    for fi60 := 0; fi60 < 10; fi60++ {
        var v61 float64 = 1e10
        _ = fi60
        _ = v61
    }
    fmt.Println(v2, v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v59
}
