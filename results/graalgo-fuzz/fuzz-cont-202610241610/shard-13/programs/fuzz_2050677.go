package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "foo bar"
    var v3 bool = false
    var v4 int = 42
    var v5 int64 = 1
    var v6 int16 = -1
    v2 = "a"
    v1[425] = false
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 = -1
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    v3 = (3.14159 < 3.14159)
    v9 := []int{(v4 & v4)}
    if true {
        if false {
            v10 := GenS1{F0: 554, F1: 42, F2: v4}
            if true {
                fmt.Println(v5, v3, v4)
                _ = len(v1)
            }
            _ = v10
        } else {
            delete(v1, 2147483648)
            var v11 int16 = v6
            _ = v11
        }
        var v12 int64 = (v5 + v5)
        v13 := GenS1{F0: 18446744073709551615, F1: 65535, F2: 4294967295}
        _ = v12
        _ = v13
    } else {
        if true {
            if v3 {
                v9 = append(v9, (v4 | (v4 | (v4 - v4))))
                v3 = v3
                v14 := GenS1{F0: 42, F1: 4294967295, F2: 32767}
                _ = v14
            }
            if (v3 || false) {
                fmt.Println(v3, v2)
                v15 := GenS2{F0: 65535}
                _ = len(v1)
                _ = v15
            }
        } else {
            if false {
                delete(v1, 65535)
            }
        }
        v16 := GenS3{F0: 255, F1: 100}
        _ = v16
    }
    if len(v9) > 0 {
        v9[0] = v4
    }
    v5 = v5
    for fi17 := 0; fi17 < 4; fi17++ {
        v2 = v2
        var v18 float32 = ((0.0 / (3.14159 + 0.1)) - 1e10)
        _ = fi17
        _ = v18
    }
    switch (("foo bar" + v2) + ("hello" + ("foo bar" + ((v2 + ((("hello" + "") + ("" + "hello")) + (v2 + (("foo bar" + v2) + ((v2 + v2) + (v2 + "a")))))) + (v2 + ("" + "\n\t")))))) {
    case "a":
        v2 = v2
    }
    if v3 {
        for fi19 := 0; fi19 < 12; fi19++ {
            if (953 >= 636) {
                if v3 {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if (((v6 - v6) < v6) || (v2 >= v2)) {
                        _ = len(v9)
                    } else {
                        var v20 uint8 = 2
                        _ = v20
                    }
                } else {
                    v4 -= fi19
                    fmt.Println(v5, v2)
                    fmt.Println(v6)
                }
            }
            _ = fi19
        }
        if true {
            v21 := GenS2{F0: 65535}
            fmt.Println(v4)
            _ = v21
        } else {
            var v22 uint16 = 2
            _ = v22
        }
        if true {
            v5 = v5
        }
    }
    var v23 int32 = -1
    var v24 uint64 = 976
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
