package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 18446744073709551615
    var v3 bool = (-32768 == -32768)
    for fi4 := 0; fi4 < 12; fi4++ {
        v3 = v3
        _ = fi4
    }
    v5 := GenS1{F0: 0}
    fmt.Println(v2, v2)
    v6 := []int{}
    v3 = v3
    v7 := GenS1{F0: 127}
    for fi8 := 0; fi8 < 2; fi8++ {
        v3 = ((100 != 119) && (("foo bar" + "\n\t") > "a"))
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS2{F0: 1}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v3 = (((1e10 + 0.0) > 3.14159) || v3)
    _ = len(v6)
    if (3.14159 <= 0.1) {
        fmt.Println(v3, v2)
        var v11 uint = 42
        _ = v11
    } else {
        v12 := GenS1{F0: 2}
        _ = len(v1)
        var v13 int32 = 454
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: 255}
    var v15 uint64 = (v2 & 2)
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (fi16 + fi16) {
        case 127:
            if v3 {
                for fi17 := 0; fi17 < 20; fi17++ {
                    _ = len(v1)
                    v15 = 498
                    _ = fi17
                }
                for fi18 := 0; fi18 < 19; fi18++ {
                    _ = len(v1)
                    fmt.Println(v2, v2, v15)
                    _ = fi18
                }
            }
        case 0:
            fmt.Println(v15)
        }
        if (-2147483648 < 364) {
            var v19 uint64 = (v2 ^ (v2 << ((753) & 63)))
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int32 = -773
                if (!v3) {
                    if v3 {
                        _ = len(v6)
                        var v22 uint16 = 2
                        _ = v22
                    } else {
                        v19 = 2
                    }
                    if v3 {
                        switch ("a" + "\n\t") {
                        case "a":
                            var v23 int = int(v2)
                            _ = v23
                        case "foo bar":
                            v21 = v21
                        }
                    } else {
                        var v24 int64 = 9223372036854775807
                        _ = v24
                    }
                } else {
                    if true {
                        var v25 int32 = v21
                        v26 := GenS2{F0: 42}
                        var v27 uint16 = 255
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
                v19 = 342
                _ = fi20
                _ = v21
            }
            var v28 int64 = 1
            _ = v19
            _ = v28
        }
        v3 = v3
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
