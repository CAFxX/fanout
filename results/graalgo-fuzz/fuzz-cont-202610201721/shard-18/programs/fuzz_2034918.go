package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 uint32 = 522
    v4 := make(map[uint32]int)
    var v5 bool = v2
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (^(fi6)) {
        case -1:
            if (4294967295 < 2) {
                var v7 uint32 = v3
                v8 := []int{}
                switch "" {
                case "foo bar":
                    var v9 uint = 18446744073709551615
                    _ = v9
                case "\n\t":
                    switch ("a" + "") {
                    case "":
                        switch (("\n\t" + ("hello" + "foo bar")) + "a") {
                        case "\n\t":
                            for fi10 := 0; fi10 < 2; fi10++ {
                                var v11 uint = 42
                                switch -2147483648 {
                                case 2147483647:
                                    var v12 int32 = 42
                                    _ = v12
                                case 32768:
                                    fmt.Println(v7, v3)
                                case -9223372036854775808:
                                    v13 := GenS1{F0: false}
                                    var v14 int8 = -1
                                    _ = v13
                                    _ = v14
                                }
                                _ = fi10
                                _ = v11
                            }
                        }
                    case "\n\t":
                        var v15 uint8 = 1
                        _ = v15
                    default:
                        var v16 uint64 = 18446744073709551615
                        _ = v16
                    }
                default:
                    if false {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = len(v8)
                    }
                }
                _ = v7
                _ = v8
            }
        case -363:
            var v17 string = ("a" + "")
            _ = v17
        default:
            v2 = (1 != 42)
        }
        _ = fi6
    }
    var v18 int16 = -209
    fmt.Println(v18, v5)
    if v5 {
        if (!(!(128 <= 666))) {
            delete(v4, 349)
            v19 := GenS1{F0: (!v5)}
            v20 := []int{523, 9223372036854775807}
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: (!(1e10 >= 1.5))}
        v5 = v5
        _ = v21
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            switch (fi22 << ((2) & 63)) {
            case 127:
                v3 = v3
            case 128:
                _ = len(v1)
            case 1:
                v2 = (!(-782 < -1))
            default:
                v1[905] = (v3 | v3)
            }
            _ = len(v1)
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 float64 = (0.1 * float64((0.1 + 3.14159)))
                v25 := []int{fi23, fi23, fi22, fi23}
                var v26 int32 = 2147483647
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
    }
    if false {
        for fi27 := 0; fi27 < 17; fi27++ {
            v5 = v5
            v28 := GenS1{F0: v2}
            switch ((("hello" + ((((("a" + "\n\t") + "\n\t") + (("\n\t" + "a") + "\n\t")) + "\n\t") + ("hello" + ""))) + "foo bar") + ("hello" + "a")) {
            case "\n\t":
                var v29 int32 = -637
                _ = v29
            default:
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 int64 = 9223372036854775807
                    v32 := []int{}
                    var v33 string = ("a" + ("\n\t" + ""))
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
            _ = fi27
            _ = v28
        }
        switch -8 {
        case 596:
            fmt.Println(v5, v5, v2)
        }
        v3 -= v3
    } else {
        var v34 float32 = 3.14159
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v5 = (!((!true) && v2))
        _ = fi35
    }
    var v36 uint16 = 0
    v37 := GenS1{F0: (v36 == 1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
