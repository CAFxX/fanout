package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 int; F1 int; F2 int16 }
func genFn0() uint64 {
    if (((!true) && false) || true) {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (!("foo bar" == "")) {
                switch "a" {
                case "hello":
                    var v8 uint = 475
                    _ = v8
                case "":
                    fmt.Println(fi7)
                }
                var v9 uint8 = 128
                _ = v9
            } else {
                var v10 bool = (!("foo bar" > "a"))
                v11 := []int{(fi7 & (-(fi7))), 996}
                var v12 uint64 = 251
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: (255 == 1)}
            if (128 > 2) {
                var v14 int32 = -2147483648
                _ = v14
            } else {
                if true {
                    var v15 float32 = (-2.25 * (0.0 / 3.14159))
                    var v16 string = ("foo bar" + "hello")
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint64 = 582
                    _ = v17
                }
            }
            _ = fi7
            _ = v13
        }
    }
    var v18 uint32 = 975
    _ = v18
    return 126
}

func main() {
    var v1 int64 = 0
    var v2 uint = 840
    var v3 int = 0
    v4 := []int{v3, 2147483647, 1, (v3 >> ((v2) & 63))}
    var v5 int32 = -2147483648
    var v6 uint64 = 90
    v19 := GenS2{F0: 255, F1: 313, F2: v5}
    var v20 uint = (v2 << ((9) & 63))
    if false {
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v6, v1, v20)
                if true {
                    fmt.Println(v5, v3)
                    var v23 int = (fi22 * (-(fi21)))
                    _ = v23
                }
                var v24 int64 = v1
                _ = fi22
                _ = v24
            }
            switch ("\n\t" + "") {
            case "a":
                v3 += (v3 >> (((v20 + v2)) & 63))
            case "foo bar":
                if (-128 == -103) {
                    switch 9223372036854775807 {
                    case 127:
                        switch "\n\t" {
                        case "a":
                            for fi25 := 0; fi25 < 11; fi25++ {
                                v5 = (v5 & 42)
                                v20 = 0
                                _ = fi25
                            }
                        case "hello":
                            _ = len(v4)
                        case "foo bar":
                            if (false && false) {
                                var v26 float64 = 3.14159
                                _ = v26
                            }
                        }
                    }
                } else {
                    v27 := GenS3{F0: v3, F1: (fi21 ^ (fi21 / (((v3 - fi21)) | 1))), F2: -690}
                    _ = v27
                }
            case "":
                var v28 int8 = 42
                _ = v28
            }
            _ = fi21
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            v3 = (v3 + (fi29 % (((v3 & (v3 / ((fi29) | 1)))) | 1)))
            _ = fi29
        }
    } else {
        var v30 uint8 = 255
        switch (v3 - (v3 % (((v3 ^ (v3 & 127))) | 1))) {
        case 9223372036854775807:
            fmt.Println(v5, v2, v5)
        default:
            for fi31 := 0; fi31 < 19; fi31++ {
                fmt.Println(v30)
                v32 := GenS2{F0: 2, F1: 578, F2: v5}
                _ = fi31
                _ = v32
            }
        }
        if true {
            if false {
                if (2 != 2147483648) {
                    var v33 int = 42
                    var v34 uint8 = (v30 ^ 255)
                    _ = v33
                    _ = v34
                }
            }
        }
        _ = v30
    }
    v3 = -1
    if false {
        fmt.Println(v5, v6)
        var v35 int32 = (v5 & -1)
        _ = v35
    } else {
        v3 -= (v3 >> ((v20) & 63))
        if true {
            var v36 float32 = (3.14159 * -2.25)
            v36 += (1.5 * v36)
            v36 += -2.25
            _ = v36
        } else {
            switch "\n\t" {
            case "hello":
                var v37 float64 = -2.25
                _ = v37
            }
            switch (v3 ^ (^(v3))) {
            case 1:
                var v38 uint32 = 783
                _ = v38
            case 765:
                var v39 uint16 = 255
                _ = v39
            }
            if ((v20 & v2) <= (v20 - v2)) {
                switch v3 {
                case 614:
                    var v40 uint16 = 99
                    _ = v40
                }
                v5 = v5
            }
        }
        var v41 uint64 = genFn0()
        _ = v41
    }
    var v42 int16 = -68
    if false {
        v5 = v5
        v5 -= (-(v5))
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
