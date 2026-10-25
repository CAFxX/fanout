package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
func genFn0(p0 uint64, p1 int, p2 bool) int {
    var v7 string = "\n\t"
    var v8 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -961
}
func genFn1(p0 int16, p1 uint16) uint16 {
    var v9 float32 = -2.25
    if false {
        if (true && (4294967295 == -2147483648)) {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int16 = (p0 ^ p0)
                _ = fi10
                _ = v11
            }
            fmt.Println(p1, p0)
            var v12 uint32 = 784
            _ = v12
        }
    }
    var v13 bool = true
    if (0 == 1) {
        var v14 uint8 = 255
        v15 := []int{}
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint16 = p1
            if (40 > 264) {
                v18 := []int{(fi16 * fi16), fi16}
                var v19 uint16 = (p1 / ((v17) | 1))
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 12; fi20++ {
                    fmt.Println(fi20, v17)
                    switch (fi16 + 906) {
                    case 32768:
                        var v21 int32 = -1
                        _ = v21
                    }
                    _ = fi20
                }
            }
            _ = fi16
            _ = v17
        }
        if v13 {
            var v22 float64 = (0.1 / 1.5)
            v13 = v13
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v9
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    var v3 bool = false
    var v4 int8 = -63
    var v5 bool = (!v3)
    v6 := []int{4294967295, 0}
    var v23 uint32 = 4294967295
    var v24 int = 127
    switch (("hello" + "") + "") {
    case "a":
        fmt.Println(v23)
    case "":
        for fi25 := 0; fi25 < 2; fi25++ {
            v24 -= fi25
            v26 := GenS1{F0: v4, F1: "a"}
            switch int(v4) {
            case 1:
                if ((((!(1 != 109)) || v3) || v3) && v5) {
                    fmt.Println(v5, v2)
                    _ = len(v1)
                }
            case -221:
                fmt.Println(fi25, v24, v24)
            }
            _ = fi25
            _ = v26
        }
    case "foo bar":
        fmt.Println(v3, v4)
    }
    v1[v23] = "foo bar"
    _ = len(v6)
    for fi27 := 0; fi27 < 16; fi27++ {
        v28 := GenS1{F0: (v4 << ((117) & 63)), F1: ("\n\t" + ("\n\t" + "foo bar"))}
        if true {
            var v29 int8 = 0
            if (!(true && v5)) {
                if (((("hello" + ("\n\t" + ("\n\t" + ((("hello" + (("" + ("\n\t" + "\n\t")) + "\n\t")) + "a") + "")))) + ("foo bar" + "a")) < (("foo bar" + "hello") + ("" + (("" + ("\n\t" + "foo bar")) + ("foo bar" + "\n\t"))))) && v5) {
                    fmt.Println(v3, v2, v29)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi30 := 0; fi30 < 5; fi30++ {
                        var v31 uint8 = 2
                        _ = fi30
                        _ = v31
                    }
                } else {
                    fmt.Println(v4, v5, v4)
                    v32 := GenS2{F0: false, F1: 702, F2: 1}
                    v33 := GenS2{F0: v5, F1: 2, F2: 0}
                    _ = v32
                    _ = v33
                }
                var v34 int = genFn0(0, (v24 % (((v24 ^ genFn0(302, int(v29), v5))) | 1)), v3)
                _ = v34
            } else {
                var v35 uint16 = 1
                _ = v35
            }
            v3 = v5
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    v5 = (132 != 1)
    var v36 bool = (65535 == 257)
    if v36 {
        if (true && (639 <= 65535)) {
            _ = len(v1)
            v36 = v3
        } else {
            if v3 {
                v23 -= (^(v23))
            }
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            var v38 int64 = 0
            var v39 int8 = (^(v4))
            _ = fi37
            _ = v38
            _ = v39
        }
        if false {
            var v40 uint = 255
            if len(v6) > 0 {
                v6[0] = (v24 & (v24 & v24))
            }
            _ = v40
        } else {
            v24 = genFn0(334, (^(v24)), v3)
        }
    }
    v23 = v23
    switch "\n\t" {
    case "hello":
        v23 += (v23 - v23)
    case "":
        if (0.1 >= 0.0) {
            fmt.Println(v2, v3)
        } else {
            var v41 string = ("a" + ("" + ("\n\t" + "a")))
            _ = v41
        }
    case "a":
        v4 = (^(v4))
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v42 int16 = 1
    var v43 float32 = 0.0
    fmt.Println(v2, v5, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
