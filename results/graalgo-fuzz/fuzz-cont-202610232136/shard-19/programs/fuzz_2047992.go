package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int32; F2 string }
type GenS6 struct { F0 uint16; F1 int32; F2 int8 }

func main() {
    var v1 bool = true
    v2 := []uint32{414, 157, 1}
    var v3 int64 = -1
    v4 := []uint32{413, 0, 203}
    for fi5 := 0; fi5 < 14; fi5++ {
        v1 = (v1 && true)
        var v6 uint = 4294967295
        _ = fi5
        _ = v6
    }
    var v7 float64 = -2.25
    v8 := GenS1{F0: "", F1: 2}
    if v1 {
        fmt.Println(v1)
    } else {
        if (!v1) {
            v3 = (v3 << ((873) & 63))
            fmt.Println(v3, v1)
            fmt.Println(v1, v3, v3)
        }
        v1 = true
    }
    v9 := GenS1{F0: (("" + ((("" + "") + "foo bar") + ("hello" + "foo bar"))) + (("a" + "\n\t") + "foo bar")), F1: 0}
    v1 = false
    v2 = append(v2, 1)
    if true {
        var v10 float32 = (0.0 - float32(0.0))
        v1 = true
        switch "\n\t" {
        case "":
            fmt.Println(v3, v3, v1)
        case "foo bar":
            v10 = v10
        case "hello":
            var v11 uint8 = 161
            _ = v11
        }
        _ = v10
    } else {
        v12 := GenS2{F0: 42}
        _ = v12
    }
    if v1 {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint32 = 0
            var v15 uint64 = 18446744073709551615
            var v16 int8 = 127
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        v7 = v7
    }
    var v17 uint64 = 529
    v3 = v3
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v17, v17)
        for fi19 := 0; fi19 < 2; fi19++ {
            if (65535 < 1) {
                v17 += (v17 ^ v17)
                v20 := GenS3{F0: v3, F1: 469}
                fmt.Println(v1)
                _ = v20
            } else {
                var v21 int = fi19
                _ = v21
            }
            switch fi18 {
            case 32768:
                _ = len(v4)
            case 42:
                v1 = (!true)
                var v22 uint = 2
                _ = v22
            case 32767:
                v1 = v1
            default:
                if false {
                    fmt.Println(v7, v7, fi18)
                    if (4294967295 < (fi18 + (fi19 + (fi19 - int(v3))))) {
                        if (672 >= 73) {
                            v23 := GenS2{F0: (v17 << ((1) & 63))}
                            _ = v23
                        }
                    }
                    if (float32(v7) != 1e10) {
                        v24 := GenS4{F0: (("hello" + "a") != "\n\t")}
                        fmt.Println(fi19)
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                        _ = v24
                    } else {
                        fmt.Println(fi19, v3, fi18)
                        var v25 uint32 = 536
                        v25 = (v25 - (v25 ^ (v25 % ((747) | 1))))
                        _ = v25
                    }
                }
                switch "hello" {
                case "":
                    switch ("\n\t" + "foo bar") {
                    case "":
                        v4 = append(v4, 448)
                    case "a":
                        v1 = false
                    default:
                        fmt.Println(v3)
                    }
                default:
                    var v26 string = ""
                    _ = v26
                }
                var v27 int64 = 478
                _ = v27
            }
            var v28 int8 = 42
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi19
            _ = v28
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 string = ""
            v31 := GenS4{F0: (125 != 1)}
            var v32 int = (fi18 - (fi29 >> ((65535) & 63)))
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi18
    }
    switch "\n\t" {
    case "hello":
        if v1 {
            v3 = v3
            v33 := GenS3{F0: (^(v3)), F1: 42}
            switch "\n\t" {
            case "foo bar":
                for fi34 := 0; fi34 < 1; fi34++ {
                    v7 += 3.14159
                    v17 = v17
                    var v35 bool = (!(false && (255 == 42)))
                    _ = fi34
                    _ = v35
                }
            case "":
                fmt.Println(v17, v3, v7)
            case "\n\t":
                if v1 {
                    if len(v4) > 0 {
                        v4[0] = 101
                    }
                    if v1 {
                        if v1 {
                            v3 = (v3 >> ((1) & 63))
                            var v36 int16 = int16(v17)
                            _ = v36
                        }
                        v7 = 3.14159
                    } else {
                        fmt.Println(v17, v7, v17)
                    }
                } else {
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v3 = v3
                        _ = fi37
                    }
                    v2 = append(v2, 4294967295)
                }
            }
            _ = v33
        }
    case "foo bar":
        if v1 {
            for fi38 := 0; fi38 < 5; fi38++ {
                for fi39 := 0; fi39 < 13; fi39++ {
                    if v1 {
                        var v40 uint16 = 2
                        v41 := GenS5{F0: 2, F1: 2147483647, F2: "\n\t"}
                        _ = v40
                        _ = v41
                    }
                    v42 := GenS6{F0: 255, F1: 42, F2: 100}
                    _ = fi39
                    _ = v42
                }
                v1 = (1 == -32768)
                var v43 int64 = (v3 - 0)
                _ = fi38
                _ = v43
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
