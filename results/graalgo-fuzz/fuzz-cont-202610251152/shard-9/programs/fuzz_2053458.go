package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 bool; F1 uint; F2 int64 }

func main() {
    var v1 int16 = 1
    var v2 uint64 = 4294967295
    var v3 int32 = -1
    v1 += (v1 ^ v1)
    var v4 uint16 = 395
    v1 = (-(v1))
    v4 = (v4 | 255)
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "" {
        case "hello":
            for fi6 := 0; fi6 < 19; fi6++ {
                v7 := []int{int(v4), fi6}
                var v8 uint16 = v4
                _ = fi6
                _ = v7
                _ = v8
            }
        case "a":
            v9 := GenS1{F0: v1}
            _ = v9
        case "\n\t":
            var v10 uint8 = 128
            _ = v10
        }
        if true {
            fmt.Println(v2)
        }
        _ = fi5
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            switch ("a" + "") {
            case "hello":
                v13 := []int{fi11, (fi11 + fi12), fi11, int(v1)}
                _ = v13
            case "a":
                var v14 int = (-(fi12))
                _ = v14
            }
            if ((255 < 231) && ((fi11 - fi12) >= 32767)) {
                switch (fi12 ^ fi11) {
                case 0:
                    fmt.Println(v4, fi11)
                }
                v4 = v4
            } else {
                v15 := GenS2{F0: v3, F1: (true || ((v3 * int32(v1)) >= v3))}
                v16 := []int{-881, fi12}
                if true {
                    switch -2147483648 {
                    case 32767:
                        if (false && (!(4294967295 == fi11))) {
                            if (220 != 648) {
                                v17 := GenS3{F0: (v2 + (v2 / ((1) | 1)))}
                                var v18 int64 = 91
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint = 1
                                switch fi11 {
                                case 9223372036854775807:
                                    var v20 uint16 = (^(v4))
                                    _ = v20
                                case 127:
                                    var v21 uint32 = 1
                                    _ = v21
                                }
                                switch (-(fi11)) {
                                case 146:
                                    v1 = (^(v1))
                                case 32768:
                                    if (false || (128 != 28)) {
                                        var v22 uint = 18446744073709551615
                                        v23 := GenS4{F0: 42}
                                        v24 := GenS4{F0: 4294967295}
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 uint64 = 2
                                        _ = v25
                                    }
                                    var v26 uint64 = v2
                                    _ = v26
                                case -1:
                                    var v27 int16 = (v1 << (((v19 + v19)) & 63))
                                    _ = v27
                                default:
                                    if (true && (!(!true))) {
                                        if (1.5 != (-2.25 / float64(1e10))) {
                                            v1 = v1
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        fmt.Println(fi11)
                                        var v28 int16 = 857
                                        if false {
                                            if true {
                                                var v29 uint = (v19 ^ v19)
                                                if (0 == 127) {
                                                    fmt.Println(fi11)
                                                    var v30 string = (("a" + "") + (("foo bar" + ("a" + ("foo bar" + "\n\t"))) + "foo bar"))
                                                    if false {
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                    } else {
                                                        _ = len(v16)
                                                        if ((3.14159 + (float64(1e10) + 0.1)) >= 3.14159) {
                                                            v29 += 33
                                                            if false {
                                                                _ = len(v16)
                                                                var v31 uint16 = 2
                                                                var v32 int8 = 100
                                                                if (3.14159 <= (-2.25 + (((-2.25 - -2.25) * float64((float32(0.0) / (1.5 + 1.5)))) / (0.1 * (1e10 / 1e10))))) {
                                                                    v30 = (v30 + "foo bar")
                                                                    v16 = append(v16, 32767)
                                                                    var v33 int32 = (v3 & (-(v3)))
                                                                    if false {
                                                                        if len(v16) >= 2 {
                                                                            v16 = v16[1:len(v16)]
                                                                        }
                                                                        v16 = append(v16, (fi12 * (-(fi11))))
                                                                        if len(v16) >= 2 {
                                                                            v16 = v16[1:len(v16)]
                                                                        }
                                                                    } else {
                                                                        v30 = ("" + v30)
                                                                        if len(v16) > 0 {
                                                                            v16[0] = 17
                                                                        }
                                                                        v29 = v19
                                                                        var v34 bool = false
                                                                        _ = v34
                                                                    }
                                                                    _ = v33
                                                                } else {
                                                                    fmt.Println(v19, fi11, v1)
                                                                }
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            switch v30 {
                                                            case "a":
                                                                switch ("a" + (v30 + (v30 + v30))) {
                                                                case "":
                                                                    if len(v16) >= 2 {
                                                                        v16 = v16[1:len(v16)]
                                                                    }
                                                                default:
                                                                    v35 := GenS5{F0: 255}
                                                                    _ = v35
                                                                }
                                                            }
                                                            v36 := GenS6{F0: 609, F1: 1}
                                                            var v37 float32 = float32(1e10)
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            fmt.Println(fi11)
                                                        }
                                                        v2 += (^(v2))
                                                    }
                                                    _ = v30
                                                } else {
                                                    v38 := GenS4{F0: 0}
                                                    _ = v38
                                                }
                                                _ = v29
                                            }
                                            var v39 uint32 = 0
                                            var v40 uint16 = (v4 + (^(v4)))
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = v28
                                    }
                                }
                                _ = v19
                            }
                            switch ("\n\t" + ((("\n\t" + "hello") + "a") + "")) {
                            case "foo bar":
                                if true {
                                    var v41 uint64 = 4294967295
                                    _ = v41
                                }
                            }
                        }
                    case -143:
                        v42 := GenS7{F0: int8(v1), F1: (-245 <= -1)}
                        _ = v42
                    case 42:
                        switch ("hello" + "\n\t") {
                        case "":
                            v3 += -2147483648
                        case "\n\t":
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        case "a":
                            fmt.Println(v4)
                        }
                    }
                    if true {
                        var v43 uint = 18446744073709551615
                        var v44 int32 = v3
                        _ = v43
                        _ = v44
                    }
                    var v45 int8 = -1
                    _ = v45
                } else {
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(fi12, v4)
                    case "a":
                        _ = len(v16)
                    }
                    _ = len(v16)
                    v46 := GenS3{F0: (^(v2))}
                    _ = v46
                }
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        _ = fi11
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        var v48 uint32 = 0
        _ = fi47
        _ = v48
    }
    if false {
        if (!(52 >= 100)) {
            v49 := GenS8{F0: true, F1: 402, F2: -243}
            var v50 float32 = 1.5
            _ = v49
            _ = v50
        }
        v1 += (v1 / ((32767) | 1))
        v3 = v3
    } else {
        v2 += v2
    }
    fmt.Println(v2)
    var v51 bool = false
    fmt.Println(v1)
    v2 = (v2 >> ((665) & 63))
    var v52 uint8 = 52
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
}
