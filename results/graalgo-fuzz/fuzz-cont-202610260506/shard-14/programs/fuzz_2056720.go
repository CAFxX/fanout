package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool, p1 bool) string {
    v4 := GenS1{F0: 2}
    v5 := GenS1{F0: 2147483647}
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, fi7}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return ("hello" + "")
}
func genFn1(p0 bool) int {
    var v9 int = -109
    v10 := GenS1{F0: 4294967295}
    _ = p0
    _ = v10
    _ = v9
    return (v9 - v9)
}

func main() {
    var v1 bool = (!(!false))
    var v2 uint32 = 968
    var v3 int64 = 42
    v1 = (!(-22 <= 104))
    var v11 int64 = (v3 + v3)
    if (!v1) {
        var v12 uint8 = 134
        if ((v2 >= v2) || v1) {
            v13 := []int{-208}
            _ = v13
        }
        if (32767 < 2) {
            var v14 float64 = (0.1 * -2.25)
            for fi15 := 0; fi15 < 6; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint = 854
                    if false {
                        if v1 {
                            switch "foo bar" {
                            case "a":
                                v11 = v3
                            case "foo bar":
                                v18 := []int{fi15, (fi16 + fi15)}
                                _ = v18
                            case "\n\t":
                                var v19 bool = ("a" >= ("\n\t" + "a"))
                                _ = v19
                            }
                            switch "a" {
                            case "a":
                                var v20 bool = v1
                                fmt.Println(fi16, v11)
                                v21 := GenS1{F0: 4294967295}
                                _ = v20
                                _ = v21
                            }
                        }
                    } else {
                        var v22 int8 = -1
                        var v23 int16 = -1
                        switch genFn1(v1) {
                        case 128:
                            fmt.Println(v12)
                        }
                        _ = v22
                        _ = v23
                    }
                    _ = fi16
                    _ = v17
                }
                var v24 uint16 = 839
                _ = fi15
                _ = v24
            }
            _ = v14
        } else {
            var v25 uint16 = 42
            var v26 uint64 = 2
            v27 := []int{42, 42, 4294967295}
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v12
    } else {
        v28 := GenS1{F0: (v2 ^ 623)}
        _ = v28
    }
    v11 = v11
    v29 := []int{}
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 int32 = -1
        if (!v1) {
            if len(v29) > 0 {
                v29[0] = -307
            }
            v32 := GenS1{F0: v2}
            for fi33 := 0; fi33 < 3; fi33++ {
                switch "\n\t" {
                case "foo bar":
                    var v34 int16 = 32767
                    _ = v34
                default:
                    var v35 int64 = (v3 + int64(fi30))
                    v31 = v31
                    var v36 int16 = -32768
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
            _ = v32
        }
        var v37 string = "a"
        _ = fi30
        _ = v31
        _ = v37
    }
    _ = len(v29)
    _ = len(v29)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
}
