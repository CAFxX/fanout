package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    var v1 bool = (0.1 >= 0.0)
    var v2 int8 = -128
    var v3 uint = 284
    var v4 int16 = 986
    v1 = ((v2 - 68) > v2)
    var v5 uint32 = 0
    fmt.Println(v2, v3, v3)
    v6 := GenS1{F0: 126}
    for fi7 := 0; fi7 < 5; fi7++ {
        if (32767 != 42) {
            var v8 uint8 = 200
            v3 = v3
            _ = v8
        }
        _ = fi7
    }
    v3 = (v3 | v3)
    if (1 < 0) {
        v9 := []int{42}
        if v1 {
            v4 = (v4 * v4)
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(v2, v4, v2)
                _ = fi10
            }
        }
        _ = v9
    } else {
        if (!(890 == 42)) {
            var v11 int16 = v4
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS2{F0: 818}
                var v14 int16 = v4
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: int32(v4)}
            _ = v11
            _ = v15
        }
        v1 = (0.0 <= 3.14159)
    }
    if v1 {
        for fi16 := 0; fi16 < 13; fi16++ {
            if (v4 < -32768) {
                var v17 int8 = (v2 * (^(v2)))
                var v18 uint8 = 1
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        v19 := GenS2{F0: uint16(v3)}
        _ = v19
    } else {
        var v20 float64 = ((1.5 + float64(float32((3.14159 - float64((-2.25 + 0.1)))))) + 1.5)
        var v21 float64 = 0.0
        _ = v20
        _ = v21
    }
    switch -1 {
    case 9223372036854775807:
        if (!(v1 || v1)) {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint = v3
                if (!v1) {
                    v23 = (v23 << (((v23 | 1)) & 63))
                    var v24 uint32 = (^(v5))
                    _ = v24
                }
                if ((v2 % ((1) | 1)) == v2) {
                    var v25 bool = false
                    var v26 uint16 = 799
                    v23 = v23
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
            v2 = (v2 + -128)
        } else {
            if (v3 == 42) {
                for fi27 := 0; fi27 < 14; fi27++ {
                    v28 := []int{(fi27 ^ (-(fi27))), (fi27 / ((-9223372036854775808) | 1)), (^(fi27)), -2147483648}
                    v29 := GenS1{F0: 62}
                    var v30 int = (-(fi27))
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            } else {
                v31 := GenS3{F0: 0, F1: 1}
                for fi32 := 0; fi32 < 20; fi32++ {
                    if ("hello" >= "") {
                        switch ("hello" + "") {
                        case "a":
                            var v33 int64 = -9223372036854775808
                            _ = v33
                        case "foo bar":
                            var v34 uint64 = 42
                            _ = v34
                        default:
                            if v1 {
                                var v35 float32 = 0.0
                                switch "hello" {
                                case "a":
                                    if v1 {
                                        var v36 float64 = 0.0
                                        var v37 uint = 65535
                                        _ = v36
                                        _ = v37
                                    }
                                case "":
                                    var v38 int = fi32
                                    _ = v38
                                case "\n\t":
                                    for fi39 := 0; fi39 < 8; fi39++ {
                                        v40 := []int{(fi39 % ((-1) | 1))}
                                        v4 = (v4 % ((0) | 1))
                                        if v1 {
                                            switch (fi39 & fi32) {
                                            case 2147483647:
                                                fmt.Println(v2, v5)
                                            default:
                                                v4 += -40
                                            }
                                        } else {
                                            v41 := GenS2{F0: 345}
                                            var v42 int = -2147483648
                                            var v43 uint = (v3 & (v3 + 904))
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                }
                                _ = v35
                            }
                        }
                    }
                    _ = fi32
                }
                _ = v31
            }
            var v44 int16 = -32768
            v45 := []int{131}
            _ = v44
            _ = v45
        }
    case 128:
        v46 := GenS1{F0: -2147483648}
        _ = v46
    }
    var v47 uint32 = 42
    fmt.Println(v3)
    v48 := []int{366, -1, 926}
    var v49 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
