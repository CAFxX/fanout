package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int64 = 0
    var v3 int16 = 0
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{(v1 - -73), v1, (-(v1))}
        _ = fi4
        _ = v5
    }
    if false {
        v3 = v3
        switch "" {
        case "hello":
            v1 = v1
        case "a":
            var v6 int8 = 42
            _ = v6
        }
    }
    var v7 int64 = 0
    if true {
        v8 := GenS1{F0: ("\n\t" + "a")}
        _ = v8
    } else {
        switch (v1 & (-(v1))) {
        case 165:
            v9 := GenS1{F0: "hello"}
            _ = v9
        case -735:
            var v10 uint8 = 2
            _ = v10
        case 1:
            if (false && (!(-2147483648 > 0))) {
                var v11 uint32 = 1
                _ = v11
            } else {
                v7 -= (v2 & (-(v7)))
                if true {
                    if true {
                        var v12 string = "\n\t"
                        for fi13 := 0; fi13 < 14; fi13++ {
                            v14 := GenS1{F0: (v12 + (v12 + v12))}
                            v12 = v12
                            _ = fi13
                            _ = v14
                        }
                        if ((1.5 * 0.0) >= 1.5) {
                            var v15 uint = 4294967295
                            var v16 string = (v12 + "")
                            v17 := GenS1{F0: v12}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = v12
                    }
                } else {
                    switch 2147483647 {
                    case -1:
                        for fi18 := 0; fi18 < 8; fi18++ {
                            for fi19 := 0; fi19 < 17; fi19++ {
                                v20 := []int{(-(fi19)), v1, fi18}
                                _ = fi19
                                _ = v20
                            }
                            _ = fi18
                        }
                    case 9223372036854775807:
                        v21 := GenS1{F0: ("" + ("a" + ""))}
                        _ = v21
                    default:
                        v22 := []int{v1, v1}
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    v2 = -1
                    v24 := GenS1{F0: ("hello" + ("foo bar" + "\n\t"))}
                    v25 := GenS1{F0: ""}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int8 = int8(v2)
            var v28 uint16 = 0
            _ = fi26
            _ = v27
            _ = v28
        }
        v2 = (-(v7))
    }
    v29 := []int{(v1 << ((211) & 63)), v1}
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS1{F0: "foo bar"}
        var v32 float32 = 0.0
        _ = fi30
        _ = v31
        _ = v32
    }
    if (!true) {
        v33 := GenS2{F0: 128, F1: 625}
        v2 += 263
        _ = v33
    }
    switch 41 {
    case 4294967295:
        var v34 uint = 0
        _ = v34
    }
    var v35 uint32 = 553
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v7
}
