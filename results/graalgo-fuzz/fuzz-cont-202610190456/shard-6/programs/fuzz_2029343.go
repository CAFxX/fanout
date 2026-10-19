package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0() uint64 {
    v7 := GenS1{F0: 65535}
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint32 = 1
            if true {
                v11 := []int{-706}
                if (("a" + "\n\t") < "hello") {
                    if len(v11) > 0 {
                        v11[0] = (fi9 - (fi9 >> ((18446744073709551615) & 63)))
                    }
                    if (0 <= 1) {
                        switch "a" {
                        case "foo bar":
                            var v12 uint8 = 2
                            _ = v12
                        case "a":
                            var v13 float64 = 0.1
                            _ = v13
                        }
                    }
                    var v14 uint64 = 2
                    _ = v14
                }
                _ = v11
            } else {
                switch ("hello" + "hello") {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        v10 = (^(v10))
                    case "foo bar":
                        v15 := GenS2{F0: 951}
                        _ = v15
                    }
                default:
                    v16 := []int{}
                    _ = v16
                }
                fmt.Println(v10, fi9, fi9)
            }
            if false {
                switch ("" + ("\n\t" + (("" + ("\n\t" + "")) + ("a" + "")))) {
                case "":
                    if (("\n\t" + "hello") <= "") {
                        switch fi8 {
                        case -2147483648:
                            v17 := GenS3{F0: 65535}
                            _ = v17
                        }
                        if true {
                            var v18 string = "foo bar"
                            _ = v18
                        } else {
                            var v19 int8 = 0
                            v20 := []int{32767}
                            switch (fi9 & 370) {
                            case 32768:
                                var v21 uint64 = 0
                                _ = v21
                            default:
                                var v22 bool = false
                                _ = v22
                            }
                            _ = v19
                            _ = v20
                        }
                    }
                }
                v10 = 2
                v23 := []int{}
                _ = v23
            } else {
                v10 = v10
                v10 -= (v10 % ((v10) | 1))
                switch "foo bar" {
                case "a":
                    v24 := GenS2{F0: (^(v10))}
                    _ = v24
                case "hello":
                    v10 = (v10 << ((2) & 63))
                case "\n\t":
                    v10 = (v10 % ((2147483647) | 1))
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        if false {
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 int8 = 36
                _ = fi26
                _ = v27
            }
            if true {
                for fi28 := 0; fi28 < 5; fi28++ {
                    switch (("hello" + (((("" + "hello") + "foo bar") + "") + "foo bar")) + ("foo bar" + "a")) {
                    case "foo bar":
                        fmt.Println(fi28)
                    case "a":
                        if true {
                            var v29 uint64 = 42
                            v30 := GenS3{F0: 591}
                            _ = v29
                            _ = v30
                        }
                    case "":
                        fmt.Println(fi25)
                    default:
                        v31 := []int{4294967295, (fi25 | (-(fi28))), 127}
                        _ = v31
                    }
                    _ = fi28
                }
                var v32 uint64 = 866
                _ = v32
            } else {
                switch 4294967295 {
                case 4294967295:
                    var v33 int16 = 0
                    _ = v33
                case -1:
                    var v34 int16 = 1000
                    _ = v34
                case 128:
                    if (!true) {
                        var v35 float64 = 0.1
                        v36 := []int{(fi25 & (^(fi25))), fi25, (fi25 ^ 0), (fi25 | (fi25 & fi25))}
                        _ = v35
                        _ = v36
                    }
                }
                var v37 uint = 255
                _ = v37
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 int16 = 0
                if true {
                    v40 := []int{2147483647, fi38, (fi38 % ((351) | 1))}
                    switch 2147483647 {
                    case 0:
                        var v41 uint = 174
                        _ = v41
                    case -853:
                        v40 = append(v40, 127)
                    case 42:
                        v39 = (v39 | v39)
                    default:
                        var v42 int16 = v39
                        _ = v42
                    }
                    v40 = append(v40, fi25)
                    _ = v40
                } else {
                    v43 := []int{672, (fi25 & fi38), 677}
                    var v44 int16 = (^(v39))
                    fmt.Println(fi25, fi38, v39)
                    _ = v43
                    _ = v44
                }
                var v45 float64 = 3.14159
                v45 += 0.1
                v46 := []int{}
                _ = fi38
                _ = v39
                _ = v45
                _ = v46
            }
        } else {
            var v47 int64 = -1
            v47 -= -1
            _ = v47
        }
        _ = fi25
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = 122
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    var v5 int = 985
    var v6 bool = ((v5 * 2147483647) != v5)
    v48 := GenS3{F0: 0}
    fmt.Println(v5, v1)
    var v49 int32 = -970
    v5 = v5
    _ = len(v4)
    v1 = -279
    var v50 uint16 = 1
    var v51 uint16 = (v50 & (v50 | v50))
    v52 := GenS2{F0: 42}
    if (v6 || true) {
        for fi53 := 0; fi53 < 18; fi53++ {
            v6 = true
            fmt.Println(v1, v6, v51)
            _ = fi53
        }
    } else {
        if v6 {
            v51 -= 344
        } else {
            v49 = 1
            var v54 int32 = int32(v50)
            v55 := GenS3{F0: v51}
            _ = v54
            _ = v55
        }
        fmt.Println(v51, v5, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
