package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint32) bool {
    v5 := GenS1{F0: p1}
    v6 := []int{403}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}
func genFn1() bool {
    var v7 uint16 = 79
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 2147483648}
        if true {
            v10 := GenS1{F0: 65535}
            if true {
                v11 := GenS1{F0: 2147483648}
                _ = v11
            } else {
                var v12 uint = 65535
                v13 := []int{}
                _ = v12
                _ = v13
            }
            var v14 uint64 = 877
            _ = v10
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    v15 := GenS1{F0: 2}
    _ = v15
    _ = v7
    return true
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 2
    var v3 int16 = 1000
    var v4 string = "\n\t"
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int = 42
        v2 -= v2
        _ = fi16
        _ = v17
    }
    if false {
        v18 := []int{922, -9223372036854775808, 0, 579}
        var v19 float64 = 0.1
        _ = v18
        _ = v19
    }
    v20 := []int{4294967295, 1, -828, 0}
    v4 = v1
    if false {
        v3 -= (^(v3))
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int32 = 2147483647
            _ = fi21
            _ = v22
        }
        switch "foo bar" {
        case "hello":
            if ((1 >= 2) && false) {
                v3 = v3
                v4 = (v1 + v4)
                switch -744 {
                case 2147483647:
                    var v23 float32 = 1.5
                    _ = v23
                case 870:
                    v24 := GenS1{F0: 65535}
                    _ = v24
                }
            }
        case "a":
            var v25 int32 = 0
            _ = v25
        case "":
            var v26 float64 = float64(1e10)
            _ = v26
        }
    }
    v4 = (v1 + v1)
    if len(v20) > 0 {
        v20[0] = 128
    }
    v27 := GenS1{F0: 65535}
    v1 = ((v1 + "\n\t") + ("a" + ""))
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 bool = (false || (!(((0.0 == (0.1 - (-2.25 + 1e10))) && ((false && false) && true)) || ("" != "hello"))))
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            v4 = (v1 + ("foo bar" + v4))
            v1 = "hello"
            _ = fi31
        }
        _ = fi30
    }
    v32 := GenS1{F0: 2147483648}
    switch (((v4 + (v1 + v4)) + ("foo bar" + v4)) + ("a" + (("a" + v1) + v1))) {
    case "\n\t":
        fmt.Println(v4, v3, v2)
    }
    if (!(!true)) {
        v3 -= -623
        v1 = v4
        for fi33 := 0; fi33 < 4; fi33++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            var v34 int32 = 1
            var v35 uint = 776
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
