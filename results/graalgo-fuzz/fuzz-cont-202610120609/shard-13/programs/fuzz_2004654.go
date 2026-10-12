package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []string{"\n\t", "\n\t"}
    var v2 bool = (65535 >= 0)
    var v3 int16 = 1000
    v4 := GenS1{F0: (4294967295 == -2147483648)}
    if false {
        if (v2 || (!v2)) {
            var v5 uint = 65535
            v6 := GenS2{F0: 0}
            _ = v5
            _ = v6
        } else {
            if (!((-1 <= -993) || true)) {
                switch 912 {
                case 127:
                    v1 = append(v1, "")
                case 32768:
                    v1 = append(v1, ("" + ((("hello" + ("a" + "\n\t")) + "\n\t") + "foo bar")))
                case 943:
                    var v7 float64 = (-2.25 - (3.14159 / 0.1))
                    _ = v7
                }
                v8 := GenS3{F0: 0}
                _ = v8
            } else {
                fmt.Println(v2)
                _ = len(v1)
            }
            v9 := GenS2{F0: 1}
            v1 = append(v1, "\n\t")
            _ = v9
        }
        switch (("hello" + "hello") + ("" + "")) {
        case "foo bar":
            if v2 {
                _ = len(v1)
                v10 := GenS1{F0: (!((!(((float64(1.5) == float64(float32((1.5 * 1.5)))) && (1 <= 2)) || (float32((0.1 / 1.0)) > 3.14159))) && (569 < 219)))}
                switch "foo bar" {
                case "":
                    var v11 float32 = 3.14159
                    _ = v11
                }
                _ = v10
            }
        case "":
            var v12 float64 = 1.5
            _ = v12
        case "\n\t":
            v1 = append(v1, ((("a" + "foo bar") + "hello") + "\n\t"))
        }
    }
    if ("" > "foo bar") {
        for fi13 := 0; fi13 < 7; fi13++ {
            if v2 {
                for fi14 := 0; fi14 < 7; fi14++ {
                    v3 -= -754
                    var v15 uint8 = 23
                    _ = fi14
                    _ = v15
                }
                var v16 uint32 = 869
                var v17 uint16 = 549
                _ = v16
                _ = v17
            }
            v1 = append(v1, "")
            _ = fi13
        }
    } else {
        var v18 int64 = 9223372036854775807
        switch 0 {
        case 30:
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS1{F0: v2}
                if (127 == -128) {
                    v21 := GenS1{F0: ((!(!v2)) || (382 != 255))}
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v2 = v2
                        _ = fi22
                    }
                    for fi23 := 0; fi23 < 2; fi23++ {
                        if v2 {
                            fmt.Println(v2, fi23, fi23)
                        } else {
                            var v24 bool = (8 < 32767)
                            _ = v24
                        }
                        _ = fi23
                    }
                    _ = v21
                } else {
                    switch (fi19 << ((825) & 63)) {
                    case 2147483647:
                        v25 := GenS3{F0: 815}
                        _ = v25
                    case 4294967295:
                        for fi26 := 0; fi26 < 16; fi26++ {
                            if (0.1 >= 0.0) {
                                if true {
                                    var v27 bool = false
                                    _ = v27
                                }
                                fmt.Println(v18, v2)
                            } else {
                                var v28 uint16 = 404
                                _ = v28
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v3 -= (-(v3))
                            fmt.Println(v3, fi26, v18)
                            _ = fi26
                        }
                    }
                }
                v29 := GenS4{F0: 103, F1: 42}
                _ = fi19
                _ = v20
                _ = v29
            }
        case -837:
            fmt.Println(v18)
        }
        _ = v18
    }
    var v30 int32 = -1
    fmt.Println(v30, v30, v3)
    v3 = v3
    var v31 uint16 = 267
    v31 -= 634
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
