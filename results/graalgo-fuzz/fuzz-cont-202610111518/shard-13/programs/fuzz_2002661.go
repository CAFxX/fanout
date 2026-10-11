package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint16 }
func genFn0(p0 int8, p1 int64, p2 int32) string {
    p0 = 127
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 42, F1: "hello", F2: "foo bar"}
        v7 := GenS1{F0: 1, F1: ("foo bar" + (((("\n\t" + "hello") + "a") + ("a" + "foo bar")) + "hello")), F2: "foo bar"}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "a"
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 100
    var v3 float32 = ((0.1 + ((-2.25 * 0.0) - 3.14159)) * 0.0)
    v4 := []uint32{4294967295, 2}
    v9 := GenS1{F0: 683, F1: "a", F2: ""}
    switch 433 {
    case 128:
        if true {
            for fi10 := 0; fi10 < 9; fi10++ {
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 float64 = ((-2.25 / 0.1) / 1e10)
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 string = "foo bar"
                    v2 = v2
                    v15 := GenS2{F0: 32767}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi10
            }
        } else {
            var v16 int64 = 42
            var v17 uint8 = 0
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint8 = (v17 << ((1) & 63))
                switch 32768 {
                case 42:
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v21 := GenS1{F0: 1000, F1: "", F2: ("hello" + "hello")}
                        v19 -= 7
                        _ = fi20
                        _ = v21
                    }
                default:
                    var v22 uint8 = 128
                    _ = v22
                }
                switch "hello" {
                case "a":
                    var v23 float32 = 3.14159
                    _ = v23
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "hello":
                    fmt.Println(fi18, v19, v1)
                }
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        }
    case -76:
        var v24 uint = 2
        _ = v24
    case 42:
        fmt.Println(v3, v1)
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        fmt.Println(v2, fi25, fi25)
        fmt.Println(fi25)
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 int = fi26
            var v28 float32 = 1.5
            switch -1 {
            case 32768:
                v28 = 0.1
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    fmt.Println(v2, v2)
    if (!(2 > 1)) {
        v29 := GenS3{F0: 42, F1: -1}
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 int = fi30
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 uint16 = 1
                v4 = append(v4, 2147483648)
                if (0 < -1) {
                    fmt.Println(v33, v33)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v34 uint = 18446744073709551615
                    _ = v34
                } else {
                    var v35 int = fi32
                    v36 := GenS1{F0: -70, F1: "a", F2: "\n\t"}
                    if false {
                        fmt.Println(v2, v35)
                    }
                    _ = v35
                    _ = v36
                }
                var v37 int32 = 42
                _ = fi32
                _ = v33
                _ = v37
            }
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        _ = len(v4)
        v4 = append(v4, 157)
        if true {
            fmt.Println(v2)
        }
    }
    if false {
        var v38 uint32 = 0
        _ = v38
    } else {
        v39 := GenS4{F0: v2}
        _ = v39
    }
    v2 = v2
    v40 := GenS5{F0: -57, F1: 2147483647, F2: 337}
    var v41 float64 = float64(v3)
    if (v3 < (3.14159 + 3.14159)) {
        switch "hello" {
        case "foo bar":
            v42 := GenS5{F0: -627, F1: 2147483647, F2: 910}
            _ = v42
        case "\n\t":
            fmt.Println(v1, v3, v1)
        default:
            if (191 < 180) {
                _ = len(v4)
            } else {
                for fi43 := 0; fi43 < 6; fi43++ {
                    var v44 bool = true
                    var v45 uint = 65535
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 5; fi46++ {
                    _ = len(v4)
                    fmt.Println(v1)
                    var v47 int = (^(fi46))
                    _ = fi46
                    _ = v47
                }
            }
        }
        var v48 int64 = v1
        var v49 int = -706
        _ = v48
        _ = v49
    }
    var v50 int8 = (v2 - (^(v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v50
    _ = v9
}
