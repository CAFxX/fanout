package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint; F1 uint16; F2 string }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 string) uint32 {
    var v4 int16 = -504
    v4 = v4
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(p0, v4, v4)
        _ = fi5
    }
    v4 -= 32767
    fmt.Println(v4, v4)
    _ = p0
    _ = v4
    return 65535
}
func genFn1(p0 int64, p1 uint16, p2 uint) int8 {
    var v6 int16 = 1
    if false {
        var v7 float64 = 0.1
        var v8 int8 = int8(p1)
        var v9 int = 42
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int = -965
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := []string{("" + ("foo bar" + "a"))}
    if true {
        var v11 uint = 0
        delete(v1, genFn0("a"))
        _ = v11
    } else {
        v12 := GenS1{F0: 32767}
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint32 = 0
        _ = fi13
        _ = v14
    }
    v3 = append(v3, "\n\t")
    v15 := GenS1{F0: 1}
    var v16 int64 = 42
    switch "hello" {
    case "a":
        if (!(true && true)) {
            v17 := GenS1{F0: 1}
            _ = v17
        }
    case "\n\t":
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := GenS1{F0: 493}
            if len(v3) > 0 {
                v3[0] = ""
            }
            _ = fi18
            _ = v19
        }
    }
    if ((-32768 >= -32768) || true) {
        delete(v1, genFn0("a"))
        var v20 float32 = float32((0.1 + (float64(3.14159) - 1.5)))
        _ = v20
    }
    if (!true) {
        v21 := GenS1{F0: 824}
        _ = v21
    } else {
        v22 := GenS1{F0: 42}
        for fi23 := 0; fi23 < 18; fi23++ {
            if true {
                v24 := GenS1{F0: 777}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 uint32 = 1
                    if (fi25 <= (^(fi25))) {
                        if ((42 < -1) || false) {
                            switch fi25 {
                            case -376:
                                var v27 bool = true
                                _ = v27
                            }
                            v28 := GenS1{F0: 1}
                            _ = v28
                        } else {
                            if false {
                                v29 := GenS2{F0: (fi23 >> ((uint(v26)) & 63))}
                                if true {
                                    v16 = -9223372036854775808
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    switch (("foo bar" + ("hello" + (("\n\t" + ("hello" + "")) + ""))) + ("hello" + (("foo bar" + "foo bar") + ("\n\t" + ("hello" + "\n\t"))))) {
                                    case "":
                                        v16 += (v16 + 0)
                                    case "a":
                                        var v30 bool = (549 == 1)
                                        _ = v30
                                    default:
                                        if (78 >= 100) {
                                            var v31 uint32 = (v26 >> ((59) & 63))
                                            v31 = (v26 >> ((255) & 63))
                                            if ((float32(1.5) / 1.5) != 1.5) {
                                                v31 = (^(v26))
                                                v32 := GenS3{F0: v31, F1: (v16 | (v16 + v16))}
                                                v33 := GenS4{F0: 733, F1: uint16(v31), F2: "a"}
                                                _ = v32
                                                _ = v33
                                            } else {
                                                var v34 int8 = 127
                                                switch "a" {
                                                case "\n\t":
                                                    v35 := GenS5{F0: 2, F1: "hello"}
                                                    _ = v35
                                                case "foo bar":
                                                    var v36 int32 = 2147483647
                                                    v3 = append(v3, "hello")
                                                    _ = v36
                                                default:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v34
                                            }
                                            v26 += 0
                                            switch (fi25 - fi23) {
                                            case 128:
                                                _ = len(v1)
                                            case 420:
                                                v37 := GenS6{F0: true, F1: 32767, F2: uint8(v26)}
                                                if (!true) {
                                                    fmt.Println(fi23, fi23)
                                                    v16 = v16
                                                } else {
                                                    v31 = v31
                                                }
                                                _ = v37
                                            }
                                            _ = v31
                                        } else {
                                            v38 := GenS7{F0: 20, F1: 0}
                                            _ = v38
                                        }
                                    }
                                    if false {
                                        var v39 int64 = 1
                                        v40 := GenS1{F0: 65535}
                                        var v41 int32 = 1
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                var v42 string = ("a" + "foo bar")
                                _ = v29
                                _ = v42
                            } else {
                                v26 = 65535
                            }
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        var v43 uint = 42
                        v43 -= (^(v43))
                        _ = v43
                    } else {
                        v2[42] = (fi23 ^ fi25)
                        if true {
                            fmt.Println(v26, fi23, v26)
                        } else {
                            if true {
                                var v44 string = ("hello" + ("\n\t" + ""))
                                _ = v44
                            }
                        }
                    }
                    _ = fi25
                    _ = v26
                }
            }
            v2[fi23] = (-(fi23))
            _ = fi23
        }
        switch "a" {
        case "a":
            for fi45 := 0; fi45 < 9; fi45++ {
                if (!(!true)) {
                    if ((-1 == -661) && (1 > -48)) {
                        v2[(fi45 % ((int(v16)) | 1))] = (fi45 + int(v16))
                        fmt.Println(v16)
                        if true {
                            var v46 int64 = (v16 ^ v16)
                            if ((1.5 * 3.14159) != 1.5) {
                                var v47 float32 = 3.14159
                                var v48 string = "hello"
                                _ = v47
                                _ = v48
                            }
                            _ = v46
                        } else {
                            if ((fi45 >= (fi45 << ((0) & 63))) || true) {
                                var v49 uint16 = 0
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v50 uint = 1
                                _ = v49
                                _ = v50
                            } else {
                                var v51 float64 = 1e10
                                fmt.Println(fi45, v16)
                                _ = v51
                            }
                        }
                    } else {
                        delete(v1, genFn0(""))
                        for fi52 := 0; fi52 < 11; fi52++ {
                            v16 = v16
                            var v53 int64 = v16
                            var v54 int = fi52
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        v55 := GenS5{F0: 32767, F1: "hello"}
                        _ = v55
                    }
                } else {
                    v16 = v16
                    fmt.Println(fi45, fi45)
                    var v56 int8 = genFn1(v16, 593, 272)
                    _ = v56
                }
                for fi57 := 0; fi57 < 19; fi57++ {
                    switch ("foo bar" + "") {
                    case "\n\t":
                        fmt.Println(fi57)
                    case "hello":
                        if true {
                            _ = len(v2)
                        }
                    case "a":
                        var v58 string = "foo bar"
                        _ = v58
                    }
                    _ = fi57
                }
                fmt.Println(v16, fi45)
                _ = fi45
            }
        case "foo bar":
            var v59 int32 = -2147483648
            _ = v59
        case "hello":
            var v60 int32 = 42
            _ = v60
        }
        _ = v22
    }
    switch "hello" {
    case "hello":
        var v61 int32 = -2147483648
        _ = v61
    case "":
        fmt.Println(v16)
    default:
        var v62 uint64 = 618
        _ = v62
    }
    var v63 int8 = 0
    fmt.Println(v63, v16)
    delete(v1, 2147483647)
    v63 = (v63 | 127)
    var v64 int64 = (v16 / ((v16) | 1))
    v65 := GenS8{F0: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v65
}
