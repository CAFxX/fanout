package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 string = ""
    if (((!(false && false)) || true) || ((0.1 < 1.5) || true)) {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint16 = 533
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: -1, F1: 275}
            var v9 int32 = 2147483647
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v10 := []int{127, -9223372036854775808, 2147483647, 792}
        _ = v10
    }
    if true {
        var v11 int64 = -1
        var v12 string = v4
        var v13 uint16 = 255
        _ = v11
        _ = v12
        _ = v13
    } else {
        if true {
            v14 := GenS2{F0: 0, F1: 0}
            if (-2147483648 >= 720) {
                v4 = v4
                if false {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 float64 = 1e10
                        _ = fi15
                        _ = v16
                    }
                }
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    v4 = v4
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 int = fi17
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(fi17, fi17)
                    _ = fi17
                }
                var v20 int64 = 9223372036854775807
                v21 := GenS3{F0: -620, F1: 42, F2: 255}
                _ = v20
                _ = v21
            }
            if (!(0.0 < float64(-2.25))) {
                v22 := GenS1{F0: 97, F1: 255}
                _ = v22
            } else {
                delete(v2, 0)
                var v23 uint = 18446744073709551615
                _ = v23
            }
            _ = v14
        }
        v4 = v4
        if (!(!false)) {
            var v24 bool = false
            fmt.Println(v24)
            v25 := GenS4{F0: -783, F1: 18446744073709551615, F2: true}
            _ = v24
            _ = v25
        } else {
            switch ("" + (((v4 + v4) + v4) + (v4 + "a"))) {
            case "a":
                v26 := GenS4{F0: 0, F1: 868, F2: true}
                _ = v26
            default:
                v4 = "foo bar"
            }
            for fi27 := 0; fi27 < 17; fi27++ {
                v4 = ("" + (v4 + v4))
                v4 = (v4 + v4)
                _ = fi27
            }
            v4 = v4
        }
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS3{F0: -97, F1: -2147483648, F2: 255}
        var v30 int32 = -2147483648
        if ((((42 != 237) && true) && false) || (42 >= 1)) {
            v30 += (v30 >> ((18446744073709551615) & 63))
            v30 = v30
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint32 = 483
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 uint64 = 4294967295
        _ = fi32
        _ = v33
    }
    var v34 int64 = 0
    var v35 uint = 255
    if true {
        v35 = 5
        if true {
            var v36 int8 = 1
            var v37 int16 = 1
            _ = v36
            _ = v37
        }
        var v38 uint16 = 42
        _ = v38
    }
    if (32767 > 598) {
        for fi39 := 0; fi39 < 6; fi39++ {
            v35 += (v35 ^ (v35 ^ (v35 * (v35 & v35))))
            v40 := GenS4{F0: 9223372036854775807, F1: v35, F2: true}
            _ = fi39
            _ = v40
        }
        v41 := []int{34, 0, 9223372036854775807}
        v31 = 65535
        _ = v41
    } else {
        var v42 uint64 = 0
        var v43 uint = (v35 >> (((v35 | (v35 - v35))) & 63))
        _ = v42
        _ = v43
    }
    fmt.Println(v35)
    fmt.Println(v35, v34, v31)
    v44 := v1[v31]
    for fi45 := 0; fi45 < 9; fi45++ {
        v31 += v44
        v35 -= 4294967295
        v44 += (v44 | 670)
        _ = fi45
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := []int{int(v44)}
        _ = fi46
        _ = v47
    }
    v48 := []int{}
    if ("\n\t" != (("a" + v4) + "\n\t")) {
        switch 0 {
        case 0:
            var v49 int = -1
            _ = v49
        }
        for fi50 := 0; fi50 < 4; fi50++ {
            delete(v2, v31)
            _ = fi50
        }
    } else {
        delete(v1, uint32(v35))
    }
    v51 := GenS1{F0: 0, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v44
    _ = v48
    _ = v51
}
