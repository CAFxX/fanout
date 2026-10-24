package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }
type GenS2 struct { F0 int64; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 string }
type GenS4 struct { F0 int8; F1 int8; F2 uint }

func main() {
    var v1 int8 = 1
    var v2 int = 127
    var v3 int16 = 32767
    var v4 uint16 = 255
    var v5 float64 = -2.25
    var v6 string = "a"
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 float32 = 1e10
            v10 := []int{-524, fi7}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: true, F1: 1, F2: 560}
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        switch v6 {
        case "foo bar":
            var v13 string = "foo bar"
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 float32 = -2.25
                _ = fi14
                _ = v15
            }
        }
        switch "a" {
        case "foo bar":
            var v16 uint = 255
            _ = v16
        case "a":
            for fi17 := 0; fi17 < 14; fi17++ {
                switch "" {
                case "a":
                    v18 := []int{(v2 % (((^(fi12))) | 1)), fi12, (v2 & v2)}
                    _ = v18
                }
                v2 = v2
                _ = fi17
            }
        case "":
            switch (v6 + ("foo bar" + "a")) {
            case "\n\t":
                var v19 uint32 = 42
                _ = v19
            }
        default:
            var v20 int32 = 1
            _ = v20
        }
        _ = fi12
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := []int{v2, (fi21 + fi21), fi21, fi21}
            _ = fi22
            _ = v23
        }
        v2 += fi21
        v24 := GenS1{F0: false, F1: 0, F2: 255}
        _ = fi21
        _ = v24
    }
    switch v2 {
    case 9223372036854775807:
        var v25 bool = (0 < 74)
        _ = v25
    case 127:
        if true {
            var v26 int32 = 42
            v3 = (v3 & v3)
            v2 = v2
            _ = v26
        }
    case 863:
        v4 = v4
    default:
        if (!false) {
            if false {
                if true {
                    var v27 uint16 = 0
                    v28 := []int{(-(v2))}
                    v29 := GenS1{F0: (v2 > (v2 - v2)), F1: 18446744073709551615, F2: 489}
                    _ = v27
                    _ = v28
                    _ = v29
                }
                v3 += v3
            }
        } else {
            v30 := GenS1{F0: (v6 != v6), F1: 2, F2: 0}
            _ = v30
        }
    }
    var v31 int = (v2 | (v2 + (v2 * 42)))
    switch ("hello" + ("\n\t" + "a")) {
    case "":
        switch v2 {
        case -2147483648:
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int = (v31 / (((v31 - (^(fi32)))) | 1))
                v34 := GenS1{F0: false, F1: 845, F2: 0}
                _ = fi32
                _ = v33
                _ = v34
            }
        default:
            var v35 int64 = 9223372036854775807
            _ = v35
        }
    }
    v36 := GenS2{F0: 422, F1: 2147483647, F2: 689}
    var v37 string = ((v6 + "hello") + (v6 + "\n\t"))
    switch (v6 + v37) {
    case "":
        switch v37 {
        case "a":
            v31 = (v31 ^ (v31 | (v31 | v31)))
        case "hello":
            if false {
                v31 += v31
                if (150 >= 0) {
                    v3 = v3
                }
            } else {
                for fi38 := 0; fi38 < 15; fi38++ {
                    v1 = v1
                    _ = fi38
                }
                if (255 <= 65535) {
                    v39 := GenS3{F0: v1, F1: false, F2: ""}
                    _ = v39
                } else {
                    fmt.Println(v6)
                    v4 += v4
                    v6 = "foo bar"
                }
            }
        case "foo bar":
            v37 = "a"
        }
    case "hello":
        var v40 int = (v31 >> ((255) & 63))
        _ = v40
    case "a":
        var v41 uint32 = 1
        _ = v41
    default:
        if false {
            for fi42 := 0; fi42 < 5; fi42++ {
                var v43 uint64 = 556
                v44 := GenS3{F0: v1, F1: false, F2: "hello"}
                _ = fi42
                _ = v43
                _ = v44
            }
            if true {
                v45 := []int{-689, v2, 4294967295, v31}
                v2 += v31
                v46 := GenS4{F0: (v1 - v1), F1: (v1 - 100), F2: 91}
                _ = v45
                _ = v46
            }
        } else {
            var v47 string = v6
            v48 := []int{v31, v2, (v2 << ((0) & 63)), v2}
            var v49 float32 = float32((-2.25 + v5))
            _ = v47
            _ = v48
            _ = v49
        }
    }
    var v50 uint = 18446744073709551615
    for fi51 := 0; fi51 < 2; fi51++ {
        if true {
            var v52 uint = (v50 << ((v50) & 63))
            _ = v52
        }
        _ = fi51
    }
    var v53 string = (v6 + ("" + ((("a" + "a") + v37) + v6)))
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v6
}
