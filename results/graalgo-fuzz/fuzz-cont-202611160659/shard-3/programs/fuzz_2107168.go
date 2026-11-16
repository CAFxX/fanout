package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint8; F1 int16 }
func genFn0(p0 int32) uint {
    p0 += p0
    v4 := GenS1{F0: 4294967295, F1: 0}
    var v5 uint32 = 2147483648
    var v6 int16 = -219
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 92
}
func genFn1() string {
    var v7 uint8 = 128
    if (true && (true || false)) {
        v7 = v7
    }
    switch 1 {
    case -2147483648:
        v8 := GenS1{F0: 4294967295, F1: 1}
        _ = v8
    case 2147483647:
        if (false && false) {
            v7 -= (^(v7))
            v9 := []int{2147483647, -578}
            _ = v9
        } else {
            if false {
                v10 := []int{}
                _ = v10
            } else {
                if true {
                    fmt.Println(v7)
                    v11 := []int{32768, 897, -946, 42}
                    _ = v11
                } else {
                    var v12 bool = (-9223372036854775808 >= -91)
                    var v13 float64 = 3.14159
                    _ = v12
                    _ = v13
                }
                var v14 uint32 = 1
                v15 := []int{}
                _ = v14
                _ = v15
            }
        }
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            v7 -= v7
            _ = fi16
        }
    }
    _ = v7
    return (("" + "hello") + "a")
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{42, 922}
    var v3 uint64 = 0
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v3 = v3
            v19 := GenS2{F0: ""}
            switch (fi17 - fi17) {
            case -1:
                if (false || false) {
                    v1 = v1
                } else {
                    v20 := GenS2{F0: "foo bar"}
                    var v21 float32 = 0.0
                    var v22 float64 = 0.1
                    _ = v20
                    _ = v21
                    _ = v22
                }
            case -2147483648:
                v3 += 42
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -469:
                v23 := GenS1{F0: 255, F1: 0}
                _ = v23
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if true {
                v24 := GenS1{F0: 809, F1: 0}
                switch (fi18 - -502) {
                case 32768:
                    _ = len(v2)
                }
                _ = v24
            } else {
                v3 += v3
                v1 = v1
                switch (fi17 ^ 768) {
                case -2147483648:
                    switch (("hello" + "foo bar") + "\n\t") {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        v25 := GenS3{F0: 1, F1: 2, F2: 42}
                        _ = v25
                    default:
                        var v26 int = fi18
                        _ = v26
                    }
                case 32768:
                    fmt.Println(fi18, fi17)
                case -274:
                    var v27 float64 = 0.0
                    _ = v27
                }
                fmt.Println(fi18, fi17)
            }
            _ = fi18
            _ = v19
        }
        var v28 uint8 = 255
        for fi29 := 0; fi29 < 13; fi29++ {
            if false {
                switch fi17 {
                case 0:
                    fmt.Println(fi17, v28, fi29)
                case 1:
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = uint32(v3)
                }
            }
            var v30 uint = 7
            _ = fi29
            _ = v30
        }
        _ = fi17
        _ = v28
    }
    var v31 uint16 = 192
    v32 := GenS4{F0: v3, F1: 201}
    v33 := GenS3{F0: 42, F1: (v3 >> ((685) & 63)), F2: 0}
    v1 = (1.5 * -2.25)
    var v34 int64 = 0
    var v35 uint8 = 127
    v36 := GenS2{F0: "\n\t"}
    v3 = v3
    if true {
        for fi37 := 0; fi37 < 14; fi37++ {
            switch "hello" {
            case "hello":
                var v38 uint = 255
                _ = v38
            case "foo bar":
                v39 := GenS2{F0: "hello"}
                _ = v39
            case "":
                _ = len(v2)
            default:
                var v40 uint32 = 757
                _ = v40
            }
            _ = fi37
        }
    } else {
        var v41 int32 = -363
        var v42 bool = (!(661 == (v41 * v41)))
        _ = v41
        _ = v42
    }
    var v43 uint = 255
    v44 := GenS3{F0: 42, F1: 583, F2: (v34 - (v34 << ((genFn0(0)) & 63)))}
    switch 755 {
    case 0:
        v1 = v1
    default:
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    v45 := GenS5{F0: (v35 & 1), F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v43
    _ = v44
    _ = v45
}
