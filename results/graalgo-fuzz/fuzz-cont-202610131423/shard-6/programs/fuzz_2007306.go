package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 string; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint32; F2 string }
type GenS6 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = ((float64(1e10) / 1.5) * 0.1)
    var v3 uint = 727
    var v4 int = 127
    if true {
        v5 := GenS1{F0: -35, F1: 0}
        for fi6 := 0; fi6 < 8; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := []int{(v4 & 32768), fi6, 32768, fi7}
                v9 := GenS1{F0: 127, F1: 1}
                v4 = v4
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        _ = v5
    }
    v10 := GenS2{F0: (^(v4)), F1: ("" + ""), F2: (-(v4))}
    switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
    case "":
        v11 := GenS3{F0: -115}
        _ = v11
    case "\n\t":
        if (2147483647 == 56) {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 int32 = int32(v4)
                v14 := []int{(^(fi12))}
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: 127, F1: 255}
            _ = v15
        }
    default:
        v16 := GenS3{F0: 126}
        _ = v16
    }
    fmt.Println(v4)
    if true {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint = (v3 >> ((42) & 63))
            _ = len(v1)
            var v19 uint8 = 7
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            delete(v1, 65535)
        }
        delete(v1, 2)
    } else {
        v20 := GenS1{F0: -1, F1: 0}
        _ = v20
    }
    if true {
        v2 = (v2 / 1.5)
    } else {
        _ = len(v1)
        var v21 uint16 = 32767
        _ = v21
    }
    var v22 uint32 = 2
    var v23 int8 = 42
    v24 := GenS3{F0: (v23 << ((2) & 63))}
    var v25 int16 = 1000
    switch "" {
    case "hello":
        var v26 string = "\n\t"
        _ = v26
    default:
        v27 := []int{v4}
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            switch (v4 ^ fi29) {
            case -414:
                v30 := GenS4{F0: (v22 | v22)}
                _ = v30
            case 374:
                v31 := GenS5{F0: 1, F1: v22, F2: ""}
                _ = v31
            default:
                v1[42] = 2147483648
                var v32 int16 = v25
                _ = v32
            }
            v33 := GenS6{F0: -1}
            v34 := GenS1{F0: v23, F1: 0}
            _ = fi29
            _ = v33
            _ = v34
        }
        var v35 int8 = v23
        _ = fi28
        _ = v35
    }
    var v36 int = int(v22)
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 bool = false
        delete(v1, 989)
        _ = fi37
        _ = v38
    }
    v39 := GenS6{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v39
    _ = v4
}
