package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int; F2 int }
type GenS4 struct { F0 int16; F1 uint8; F2 bool }
type GenS5 struct { F0 string; F1 int8 }
func genFn0(p0 bool, p1 int64) int16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            switch (("\n\t" + "hello") + "hello") {
            case "":
                switch ("hello" + "") {
                case "a":
                    p0 = (-510 <= fi6)
                }
                switch fi6 {
                case 698:
                    fmt.Println(fi7, fi7)
                case 753:
                    p1 -= -9223372036854775808
                case 32768:
                    v8 := GenS1{F0: 614, F1: 1}
                    _ = v8
                }
            default:
                if (42 != 42) {
                    v9 := []int{(-(fi6)), 428, 32767, fi6}
                    _ = v9
                }
            }
            if (!true) {
                v10 := GenS1{F0: 4294967295, F1: 42}
                _ = v10
            } else {
                var v11 uint = 18446744073709551615
                var v12 float64 = 1e10
                var v13 string = "hello"
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi7
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint32 = 1
            v15 -= v15
            _ = fi14
            _ = v15
        }
        _ = fi6
    }
    var v16 bool = (!(158 > 250))
    v17 := []int{0, 128, 128}
    v18 := GenS2{F0: true}
    var v19 uint16 = 60
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return -1
}
func genFn1() int8 {
    var v20 uint8 = 158
    if (float64(((1.5 * (3.14159 * float32(float64(1.5)))) / 3.14159)) > 0.1) {
        v21 := []int{}
        if false {
            var v22 int8 = -26
            _ = v22
        }
        _ = v21
    } else {
        var v23 uint8 = v20
        if (!(42 <= 948)) {
            switch 42 {
            case 1:
                v24 := []int{9223372036854775807, 4294967295, 32768, 127}
                _ = v24
            case 32767:
                fmt.Println(v20)
            }
        } else {
            v20 = (v23 * v23)
            if true {
                if (true || (!false)) {
                    v25 := GenS3{F0: 255, F1: 2147483647, F2: -185}
                    for fi26 := 0; fi26 < 15; fi26++ {
                        switch (fi26 >> ((45) & 63)) {
                        case -2147483648:
                            var v27 uint16 = 659
                            _ = v27
                        case 4294967295:
                            switch fi26 {
                            case 127:
                                v28 := []int{-387, fi26, (fi26 << ((255) & 63))}
                                _ = v28
                            case 741:
                                v29 := []int{}
                                _ = v29
                            case 42:
                                var v30 uint16 = 869
                                _ = v30
                            }
                        case 128:
                            var v31 uint32 = 997
                            _ = v31
                        }
                        switch "a" {
                        case "a":
                            fmt.Println(v23)
                        case "\n\t":
                            var v32 int32 = -1
                            _ = v32
                        case "hello":
                            fmt.Println(fi26, v23, v20)
                        default:
                            v33 := GenS4{F0: genFn0((!(42 == 0)), -314), F1: 158, F2: false}
                            _ = v33
                        }
                        v20 += (v20 << ((417) & 63))
                        _ = fi26
                    }
                    v20 = 115
                    _ = v25
                }
                for fi34 := 0; fi34 < 15; fi34++ {
                    v23 = v23
                    if false {
                        v23 = (v23 << ((816) & 63))
                        v35 := GenS1{F0: 2, F1: 974}
                        var v36 string = "\n\t"
                        _ = v35
                        _ = v36
                    }
                    _ = fi34
                }
                var v37 int8 = 127
                _ = v37
            } else {
                var v38 int = 32768
                _ = v38
            }
        }
        if false {
            var v39 uint64 = 1
            v39 += (v39 | v39)
            v20 = (v23 % ((0) | 1))
            _ = v39
        }
        _ = v23
    }
    _ = v20
    return -1
}

func main() {
    var v1 string = ("\n\t" + ("a" + "hello"))
    var v2 uint64 = 2
    var v3 int64 = 537
    var v4 bool = (!(!false))
    v5 := make(map[string]bool)
    fmt.Println(v3)
    for fi40 := 0; fi40 < 6; fi40++ {
        switch ((v1 + v1) + v1) {
        case "foo bar":
            v5[("" + (v1 + (("\n\t" + ("a" + "\n\t")) + "a")))] = ((1 != 2) || (42 == (v3 & (v3 & (v3 >> ((255) & 63))))))
        }
        for fi41 := 0; fi41 < 5; fi41++ {
            if true {
                switch v1 {
                case "foo bar":
                    v42 := GenS5{F0: "a", F1: 120}
                    _ = v42
                case "hello":
                    switch fi41 {
                    case 2147483647:
                        v1 = v1
                    case -9223372036854775808:
                        v2 = 4294967295
                    case 32767:
                        v3 = v3
                    default:
                        if v4 {
                            var v43 uint64 = (v2 & v2)
                            var v44 int16 = genFn0(((!(!(1 != 0))) && v4), v3)
                            _ = v43
                            _ = v44
                        } else {
                            v5[v1] = ((v1 + v1) == ((v1 + (v1 + "")) + v1))
                            fmt.Println(fi41)
                        }
                    }
                case "":
                    fmt.Println(v2)
                }
            }
            _ = fi41
        }
        _ = fi40
    }
    var v45 uint8 = 0
    if (42 <= -1) {
        for fi46 := 0; fi46 < 19; fi46++ {
            switch (fi46 & 4294967295) {
            case -9223372036854775808:
                v47 := GenS3{F0: 0, F1: fi46, F2: fi46}
                _ = v47
            case 0:
                var v48 string = ""
                _ = v48
            case 32767:
                var v49 uint64 = (v2 & (v2 + v2))
                _ = v49
            default:
                var v50 uint = 0
                _ = v50
            }
            switch 4294967295 {
            case 127:
                v1 = (("" + (v1 + "\n\t")) + (("" + v1) + v1))
            case 4294967295:
                var v51 int32 = 1
                _ = v51
            default:
                if (!(!((!(v4 && (1 >= 1))) || v4))) {
                    for fi52 := 0; fi52 < 6; fi52++ {
                        if (!true) {
                            fmt.Println(v1, v45)
                            if v4 {
                                v4 = v4
                            }
                        }
                        _ = fi52
                    }
                    v53 := v5[v1]
                    for fi54 := 0; fi54 < 9; fi54++ {
                        v3 = (v3 / (((v3 >> ((0) & 63))) | 1))
                        var v55 int64 = (v3 | 810)
                        _ = fi54
                        _ = v55
                    }
                    _ = v53
                }
            }
            switch v1 {
            case "foo bar":
                fmt.Println(fi46)
            case "\n\t":
                for fi56 := 0; fi56 < 17; fi56++ {
                    v57 := []int{fi46, 486}
                    _ = fi56
                    _ = v57
                }
            case "hello":
                v58 := GenS1{F0: v2, F1: 42}
                _ = v58
            default:
                v45 = v45
            }
            _ = fi46
        }
    } else {
        v4 = true
    }
    var v59 uint32 = 0
    var v60 float32 = float32((3.14159 * (3.14159 * 1.5)))
    v59 = v59
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v59
    _ = v60
}
