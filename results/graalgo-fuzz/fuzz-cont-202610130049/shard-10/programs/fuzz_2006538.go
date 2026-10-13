package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 bool = false
    var v2 uint32 = 4294967295
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 int8 = 127
    for fi6 := 0; fi6 < 9; fi6++ {
        v5 = v5
        _ = fi6
    }
    switch ("\n\t" + "") {
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint64 = 2
            switch ("foo bar" + "") {
            case "hello":
                var v9 int = fi7
                _ = v9
            default:
                for fi10 := 0; fi10 < 4; fi10++ {
                    if ((fi10 << ((2) & 63)) <= fi7) {
                        v11 := GenS1{F0: 997}
                        _ = v11
                    }
                    v5 -= 1
                    if v1 {
                        v1 = true
                        v3 = (^(v3))
                        v3 = v3
                    } else {
                        if v1 {
                            var v12 int64 = 9223372036854775807
                            var v13 int8 = v5
                            _ = v12
                            _ = v13
                        }
                        v14 := []int{fi10}
                        switch fi10 {
                        case -464:
                            v3 -= 0
                        default:
                            switch (-(fi10)) {
                            case 128:
                                _ = len(v14)
                            case -937:
                                var v15 int8 = v5
                                _ = v15
                            case 680:
                                v16 := GenS2{F0: 917}
                                _ = v16
                            }
                        }
                        _ = v14
                    }
                    var v17 int = fi10
                    _ = fi10
                    _ = v17
                }
            }
            _ = fi7
            _ = v8
        }
    case "a":
        v4["\n\t"] = 42
    case "foo bar":
        for fi18 := 0; fi18 < 14; fi18++ {
            if (1e10 != 0.0) {
                var v19 uint64 = 567
                var v20 int = 616
                _ = v19
                _ = v20
            }
            var v21 int32 = int32(v3)
            delete(v4, "foo bar")
            _ = fi18
            _ = v21
        }
    default:
        var v22 string = "foo bar"
        _ = v22
    }
    var v23 uint = 954
    fmt.Println(v2)
    fmt.Println(v1)
    _ = len(v4)
    var v24 uint64 = 2
    for fi25 := 0; fi25 < 10; fi25++ {
        switch "a" {
        case "foo bar":
            v26 := []int{fi25, 32768, (fi25 >> (((^(v23))) & 63)), 9223372036854775807}
            _ = v26
        case "a":
            var v27 int8 = 100
            _ = v27
        default:
            var v28 uint8 = 255
            _ = v28
        }
        var v29 uint64 = v24
        var v30 float32 = 0.1
        _ = fi25
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: 936}
    for fi32 := 0; fi32 < 4; fi32++ {
        if true {
            switch ("hello" + "hello") {
            case "foo bar":
                v3 = (v3 / ((uint16(v2)) | 1))
            case "a":
                delete(v4, ("a" + "a"))
            default:
                v23 -= 42
            }
        }
        v2 = (v2 / ((v2) | 1))
        _ = fi32
    }
    v33 := []int{127}
    v23 += v23
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
}
